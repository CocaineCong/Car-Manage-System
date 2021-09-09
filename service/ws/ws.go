package ws

import (
	"CarDemo1/conf"
	"CarDemo1/service/ws/e"
	. "CarDemo1/service/ws/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

const month = 60 * 60 * 24 * 30

type SendMsg struct {
	Type		int		`json:"type"`
	Content		string	`json:"content"`
}
//{"type": 1, "content:" "abab"}
// Client 用户类
type Client struct {
	ID		string
	SendID	string
	Socket	*websocket.Conn
	Send	chan []byte
}

// Broadcast 广播类，包含广播内容和源用户（用于进行反馈）
type Broadcast struct {
	Client		*Client
	Message		[]byte
	Type		int
}

// ClientManager 用户管理
type ClientManager struct {
	Clients		map[string]*Client
	Broadcast	chan *Broadcast
	Reply		chan *Client
	Register	chan *Client
	Unregister	chan *Client
}

// Message 信息转json（包括：发送者、接受者、内容） 改进：添加头像的url，对发送者身份进行验证（token验证）
type Message struct {
	Sender		string		`json:"sender,omitempty"`
	Recipient	string		`json:"recipient,omitempty"`
	Content		string		`json:"content,omitempty"`
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client),		// 参加连接的用户，出于性能考虑，需要设置最大连接数
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:		make(chan *Client),
	Unregister: make(chan *Client),
}

// Start 项目运行前, 协程开启start -> go Manager.Start()
func (manager *ClientManager) Start() {
	for {
		log.Println("<---监听管道通信--->")
		select {
			// 建立连接
			case conn := <-Manager.Register:
				log.Println("建立新连接:%v", conn.ID)
				Manager.Clients[conn.ID] = conn		// 将此连接加入到管理器中，ID分类
				//jsonMessage, _ := json.Marshal(&Message{Content: "已连接至服务器"})
				//conn.Send <- jsonMessage
				//conn.Send <- []byte("已连接至服务器")
				message := "已连接至服务器"
				_ = conn.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"content\": \"%s\", \"code\": %d}", message, e.WebsocketSuccess)))
			// 断开连接
			case conn := <-Manager.Unregister:
				log.Println("连接关闭:%v", conn.ID)
				if _, ok := Manager.Clients[conn.ID]; ok {
					//jsonMessage, _ := json.Marshal(&Message{Content: "连接已断开"})
					//conn.Send <- jsonMessage
					//conn.Send <- []byte("连接已断开")
					message := "连接已断开"
					_ = conn.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"content\": \"%s\", \"code\": %d}", message, e.WebsocketEnd)))
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			// 广播信息
			case broadcast := <-Manager.Broadcast:
				message := broadcast.Message
				sendId	:= broadcast.Client.SendID
				flag := false			// 默认对方不在线
				for id, conn := range Manager.Clients {
					// ToDo: 出现未知问题，只能向最新的连接发送广播--也不是不行，也挺好的
					if id != sendId{
						continue
					}
					select {
						case conn.Send <- message:
							// 向对方发送信息成功
							flag = true
						default:
							// 收不到消息就关了他
							close(conn.Send)
							delete(Manager.Clients, conn.ID)
						}
				}
				id := broadcast.Client.ID
				if flag {
					log.Println("对方在线应答")
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"content\": \"对方在线应答\", \"code\": %d}", e.WebsocketOnlineReply)))
					err := InsertOne(conf.MongoDBName, id, string(message), 1, int64(month * 3))
					if err != nil {
						fmt.Println("FUCK")
						fmt.Println(err)
					}
				}else {
					log.Println("对方不在线应答")
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"content\": \"对方不在线应答\", \"code\": %d}", e.WebsocketOfflineReply)))
					err := InsertOne(conf.MongoDBName, id, string(message), 0, int64(month * 3))
					fmt.Println("FXXK")
					fmt.Println(err)
				}
		}
	}
}

// creatId 生成房间号。相当于把聊天室改为单向发送广播
func creatId(uid,touid string) string {
	return uid+"->"+touid
}

// Read connect监听用户消息，并向指定房间广播
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	//c.Socket.ReadJSON()
	for {
		c.Socket.PongHandler()
		//_, message, err := c.Socket.ReadMessage()
		sendMsg := new(SendMsg)
		//_, m, _ := c.Socket.ReadMessage()
		//fmt.Println(m)
		err := c.Socket.ReadJSON(&sendMsg)
		//fmt.Println(sendMsg)
		if err != nil {
			// 连接失败叫客户端关闭连接
			//fmt.Println(err)
			log.Println(err)
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		if sendMsg.Type == 1 {
			r1, _ := conf.RedisClient.Get(c.ID).Result()
			r2, _ := conf.RedisClient.Get(c.SendID).Result()
			if r1 >= "3" && r2 == ""{
				_ = c.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"content\": \"达到限制\", \"code\": %d}", e.WebsocketLimit)))
				_, _ = conf.RedisClient.Expire(c.ID, time.Hour*24*30).Result() // 防止重复骚扰，未建立连接刷新过期时间一个月
				continue
			} else {
				conf.RedisClient.Incr(c.ID)
				_, _ = conf.RedisClient.Expire(c.ID, time.Hour*24*30*3).Result() // 防止过快“分手”，建立连接三个月过期
			}
			log.Println(c.ID, "发送信息:", sendMsg.Content)
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMsg.Content),
			}	// 将读取到的信息，直接进行广播操作，如果对方在线，则保存为已读信息（定时过期），如果不在线，则保存为未读信息。（Redis）
		}else if sendMsg.Type == 2 {   // msg 时间戳
			// 拉取历史记录、
			time, err := strconv.Atoi(sendMsg.Content)
			if err != nil {
				time = 9999999999
			}
			//fmt.Println(time)
			results, _ := FindMany(conf.MongoDBName, c.SendID, c.ID, int64(time), 10)
			if len(results) > 10 {
				results = results[:10]
			}else if len(results) == 0{
				// 没有信息
				_ = c.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"content\": \"到底了\", \"code\": %d}", e.WebsocketEnd)))
				continue
			}
			for _, result := range results {
				_ = c.Socket.WriteMessage(websocket.TextMessage, []byte(
					"{\"msg\": "+result.Msg+", \"from\": \""+result.From+"\"}"))
			}
		}else if sendMsg.Type == 3 {   //msg == nil
			// 第一次请求
			results, err := FirstFind(conf.MongoDBName, c.SendID, c.ID)
			if err != nil {
				log.Println(err)
			}
			for _, result := range results {
				_ = c.Socket.WriteMessage(websocket.TextMessage, []byte(
					"{\"msg\": "+result.Msg+", \"from\": \""+result.From+"\"}"))
			}
		}
	}
}

// Write connect接收广播并展示给用户
func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
			case message, ok := <-c.Send:
				if !ok {
					_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				log.Println(c.ID, "接收信息:", string(message))

				_ = c.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"content\": \"%s\", \"code\": %d}", message, e.WebsocketSuccessMessage)))
			}
	}
}

//WsHandler socket 连接 中间件 作用:升级协议,用户验证,自定义信息等
func WsHandler(c *gin.Context) {
	uid := c.Query("uid")   		//token解析的ID
	touid := c.Query("to_uid")  	//聊天的对方的ID
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//可以添加用户信息验证
	client := &Client{
		ID:    creatId(uid,touid),
		SendID:creatId(touid,uid),
		Socket: conn,
		Send:   make(chan []byte),
	}
	Manager.Register <- client
	go client.Read()
	go client.Write()
}