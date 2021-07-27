package e
const (
	Success               = 200
	UpdatePasswordSuccess = 201
	NotExistInentifier    = 202
	Error                 = 500
	InvalidParams         = 400



	//成员错误
	ErrorExistNick          = 10001
	ErrorExistUser          = 10002
	ErrorNotExistUser       = 10003
	ErrorNotCompare         = 10004
	ErrorNotComparePassword = 10005
	ErrorFailEncryption     = 10006
	ErrorNotExistProduct    = 10007
	ErrorNotExistAddress    = 10008
	ErrorExistFavorite      = 10009
	//车的错误
	ErrorExistCar             = 20001
	ErrorUpdatePhone = 20002
	ErrorUpdateEmail = 20003

	//Code错误
	ErrorCodeReq = 20004
	ErrorCodeResp = 20005
	ErrorCodeOrder = 20006

	//发送短信错误
	ErrorSendMsg = 20007
	ErrorMsgCode = 20008

	//上传错误
	ErrorUploadFile = 20009

	//评论错误
	ErrorCommentNotFound = 20010
	ErrorCommentError = 20011

	//亲友错误
	ErrorFriendFound = 20012
	ErrorCreateFriend = 20013


	//管理员错误
	ErrorAuthCheckTokenFail        = 30001         //token 错误
	ErrorAuthCheckTokenTimeout     = 30002         //token 过期
	ErrorAuthToken                 = 30003
	ErrorAuth                      = 30004
	ErrorAuthInsufficientAuthority = 30005
	ErrorReadFile                  = 30006
	ErrorSendEmail                 = 30007
	ErrorCallApi                   = 30008
	ErrorUnmarshalJson             = 30009
	ErrorAdminFindUser             = 30010
	//数据库错误
	ErrorDatabase = 40001

	//通信信号
	WebsocketSuccessMessage  = 50001
	WebsocketSuccess = 50002
	WebsocketOnlineReply = 50003
	WebsocketOfflineReply = 50004
	WebsocketLimit = 50005

	//返回历史纪录
	WebsocketMsg = 60001
	WebsocketRead = 60002
	WebsocketUnread = 60003


)