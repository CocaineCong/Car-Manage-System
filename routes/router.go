package routes

import (
	"CarDemo1/api"
	"CarDemo1/conf"
	"CarDemo1/middleware"
	"CarDemo1/service/ws"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	conf.Init()
	r:=gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	//r.Use(middleware.Cors())
	middleware.HttpLogToFile(conf.AppMode)
	r.Use(middleware.LoggerToFile())
	r.Use(sessions.Sessions("mysession",store))
	v1 := r.Group("api/v1")
	{
		v1.GET("/MessageIndex/:id",ws.MessageIndex) 		//获取最新信息列表
		v1.GET("/user/get-code",api.UserGetCode)  	    //获得code,绑定手机
		v1.POST("/user/login",api.UserLogin)	  	  		//用户登陆
		v1.GET("/get-topic",api.GetTopic)			    //获取全部话题
		v1.GET("/get-social",api.GetSocial)			    //获取全部帖子
		v1.POST("/upload",api.UpLoad)
		v1.GET("/ws", ws.WsHandler)  						//通信
		v1.GET("/get-user-id/:id",api.MessageUserInfo)		//获取聊天好友的信息
		CommentGroup := v1.Group("/comment/")
		{
			CommentGroup.GET("get-single/:id", api.ShowSingleComm)
			CommentGroup.GET("get-children", api.ShowSingleChildren)
			CommentGroup.GET("get-all/:id", api.ShowAllComment)
			CommentGroup.GET("children/index", api.ShowAllComChildren)
			CommentGroup.POST("create-comment", api.CreateComment)
		}
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("ping",api.CheckToken)  				   		//验证token
			authed.GET("get-my-friend/:id",api.ShowMyFriend)  		//获取我的好友
			authed.GET("show-my-friend/:id",api.ShowMyFriendInfo)  //展示我的好友的信息
			authed.POST("create-my-friend/:id",api.CreateFriend)   //关注好友
			authed.POST("delete-my-friend/:id",api.DeleteFriend)   //删除好友
			authed.GET("user/show",api.UserShow)		  	    //获取用户信息
			authed.POST("user/vaild-email", api.VaildEmail)    //绑定邮箱
			authed.POST("user/vaild-phone",api.VaildPhone)     //绑定手机
			//authed.GET("/get-user-id/:id",api.MessageUserInfo)//

			authed.POST("create-social/:content",api.CreateSocial) 	//创建帖子
			authed.POST("search-social",api.SearchSocial)  			//搜索帖子
			authed.DELETE("social/:id", api.DeleteSocial)  			//删除帖子
			authed.GET("get-my-social/:id",api.GetMySocial)  			//获得我的帖子
			authed.GET("get-detail/:id",api.ShowSocial)	  			//获取详细的帖子

			authed.GET("cars", api.ShowCar)					//展示车
			authed.POST("cars", api.CreateCar)					//绑定车
			authed.GET("car/:id", api.DeleteCar)				//解绑车
			authed.POST("car", api.SearchCar)					//搜索车

			authed.POST("report",api.CreateReport)      			//创造Report
			authed.GET("delete-my-report/:id",api.DeleteReport)     //删除Report
			authed.GET("get-my-report",api.ShowReport)   		//获取Report

			//authed.POST("upload",api.UpLoad)							//上传操作
			//authed.POST("user/sending-email", api.SendEmail) 			//邮箱发送
			authed.POST("comment/delete-comment", api.DeleteComment)   //删除评论
			//authed.GET("/ws", ws.WsHandler)  //通信

			//authed.GET("MessageIndex/:id",ws.MessageIndex) 			//获取最新信息列表
		}
	}
	v2 := r.Group("api/v2")
	{
		//管理员登陆注册
		v2.POST("admin/register", api.AdminRegister)
		//登陆
		v2.POST("admin/login", api.AdminLogin)
		v2.GET("topic", api.GetTopic)   			//获得分类
		v2.GET("socials", api.GetSocial)  			//获得帖子
		v2.GET("socials/:id", api.ShowSocial)  		//获取帖子
		v2.GET("carousels", api.ListCarousels)  	//获得轮播图
		v2.GET("users", api.ListUsers)   			//获取用户列表
		v2.GET("reports", api.ShowReport)   		//获取用户列表
		v2.GET("cars",api.ListCars)					//获取车辆
		v2.DELETE("cars/:car_num",api.DeleteCar)					//获取车辆
		authed2 := v2.Group("/")
		authed2.Use(middleware.JWTAdmin())
		{
			authed2.POST("carousels", api.CreateCarousel)  		//创建轮播图
			authed2.DELETE("products/:id", api.DeleteSocial) 	//删除帖子
			authed2.POST("/update-social",api.UpdateSocial) 	//更新帖子
			authed2.POST("categories", api.CreateCategory)  	//创建分类
			authed2.PUT("report",api.UpdateReport)   			//更新Report
		}
	}
	return r
}