package conf

import (
	"CarDemo1/model"
	"CarDemo1/pkg/logging"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo" //MongoDB的Go驱动包
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
)


var (
	RedisClient *redis.Client		//RedisClient Redis缓存客户端单例
    MongoDBClient 		*mongo.Client
	AppMode  			string
	HttpPort 			string
	Db         			string
	DbHost     			string
	DbPort     			string
	DbUser     			string
	DbPassWord 			string
	DbName     			string

	RedisDb    			string
	RedisAddr  			string
	RedisPw    			string
	RedisDbName    		string

	MongoDBName    		string
	MongoDBAddr  		string
	MongoDBPwd    		string
	MongoDBPort    		string


	AppID 				string
	Secret 				string

	TxSecretId			string
	TxSecretKey			string
	TxSmsSign	  		string
	TxSmsSdkAppid  		string
	TxTemplateID		string

	AccessKey      string
	SerectKey      string
	Bucket     	string
	QiniuServer      string



)

func Init() {
	//从本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	LoadRedisData(file)
	LoadWxChat(file)
	LoadTxSms(file)
	LoadQiniu(file)
	LoadMongoDB(file)
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logging.Info(err) //日志内容
		panic(err)
	}
	//MySQL
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
	//MongoDB
	MongoDB()
	Redis()
}

func MongoDB()  {
	// 设置mongoDB客户端连接信息
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://"+MongoDBAddr+":"+MongoDBPort)
	// Connect to MongoDB
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logging.Info(err)
	}
	// Check the connection
	err = MongoDBClient.Ping(context.TODO(), nil)
	if err != nil {
		logging.Info(err)
	}
	logging.Info("MongoDB Connect")
}


//Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64) 		//TODO 这里记得了！！
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		//Password: conf.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}


func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").MustString("debug")
	HttpPort = file.Section("service").Key("HttpPort").MustString(":3000")
}


func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").MustString("mysql")
	DbHost = file.Section("mysql").Key("DbHost").MustString("localhost")
	DbPort = file.Section("mysql").Key("DbPort").MustString("3306")
	DbUser = file.Section("mysql").Key("DbUser").MustString("root")
	DbPassWord = file.Section("mysql").Key("DbPassWord").MustString("root")
	DbName = file.Section("mysql").Key("DbName").MustString("carsys")
}

func LoadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").MustString("redis")
	RedisAddr = file.Section("redis").Key("RedisAddr").MustString("127.0.0.1:6379")
	RedisPw = file.Section("redis").Key("RedisPw").MustString("root")
	RedisDbName = file.Section("redis").Key("RedisDbName").MustString("2")
}

func LoadWxChat(file *ini.File) {
	AppID = file.Section("wechat").Key("APPID").String()
	Secret = file.Section("wechat").Key("SECRET").String()
}

func LoadTxSms(file *ini.File) {
	TxSecretId = file.Section("txsms").Key("SecretId").String()
	TxSecretKey = file.Section("txsms").Key("SecretKey").String()
	TxSmsSign = file.Section("txsms").Key("TxSmsSign").String()
	TxSmsSdkAppid = file.Section("txsms").Key("TxSmsSdkAppid").String()
	TxTemplateID = file.Section("txsms").Key("TxTemplateID").String()
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SerectKey = file.Section("qiniu").Key("SerectKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}

func LoadMongoDB(file *ini.File) {
	MongoDBName = file.Section("MongoDB").Key("MongoDBName").MustString("userV1.0")
	MongoDBAddr = file.Section("MongoDB").Key("MongoDBAddr").MustString("localhost")
	MongoDBPwd = file.Section("MongoDB").Key("MongoDBPwd").MustString("root")
	MongoDBPort = file.Section("MongoDB").Key("MongoDBPort").MustString("27017")
}