package e

var MsgFlags = map[int]string{
	Success:               "ok",
	UpdatePasswordSuccess: "修改密码成功",

	NotExistInentifier:    "该第三方账号未绑定",
	Error:                 "fail",
	InvalidParams:         "请求参数错误",
	//用户

	ErrorExistNick:          "已存在该昵称",
	ErrorExistUser:          "已存在该用户名",
	ErrorNotExistUser:       "该用户不存在",
	ErrorNotCompare:         "账号密码错误",
	ErrorNotComparePassword: "两次密码输入不一致",
	ErrorFailEncryption:     "加密失败",

	//车
	ErrorExistCar :           "车不存在",
	ErrorUpdatePhone :"绑定手机失败",
	ErrorUpdateEmail :"绑定邮箱失败",
	ErrorCarNotFound :"该车主人未绑定",
	//Code
	ErrorCodeReq 	: "Code请求错误",
	ErrorCodeResp 	: "Code响应错误",
	ErrorCodeOrder  : "Code其他错误",

	//SendMsg
	ErrorSendMsg : "发送信息错误",
	ErrorMsgCode : "验证码错误",

	//上传错误
	ErrorUploadFile :"上传错误",

	//评论错误
	ErrorCommentNotFound :"评论不存在",
	ErrorCommentError : "评论错误",

	ErrorFriendFound : "亲友错误",
	ErrorCreateFriend :"创造亲友错误",

	ErrorAuthCheckTokenFail:        "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:     "Token已超时",
	ErrorAuthToken:                 "Token生成失败",
	ErrorAuth:                      "Token错误",
	ErrorAuthInsufficientAuthority: "权限不足",
	ErrorReadFile:                  "读文件失败",
	ErrorSendEmail:                 "发送邮件失败",
	ErrorCallApi:                   "调用接口失败",
	ErrorUnmarshalJson:             "解码JSON失败",

	//管理员
	ErrorAdminFindUser:             "管理员查询用户失败",
	ErrorDatabase: 					"数据库操作出错,请重试",

	//聊天
	WebsocketSuccessMessage : "解析content内容信息",
	WebsocketSuccess :"发送信息，请求历史纪录操作成功",
	WebsocketOnlineReply :"针对回复信息在线应答成功",
	WebsocketOfflineReply :"针对回复信息离线回答成功",
	WebsocketLimit :"请求收到限制",

	//消息记录
	WebsocketMsg :"历史纪录-对方消息",
	WebsocketRead :"历史纪录-已读消息",
	WebsocketUnread :"历史纪录-未读消息",

}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}