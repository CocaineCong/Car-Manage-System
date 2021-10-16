package service

import (
	"CarDemo1/cache"
	"CarDemo1/conf"
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	logging "github.com/sirupsen/logrus"
	"CarDemo1/pkg/util"
	"CarDemo1/serializer"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711" //引入sms
	"math/rand"
	"strconv"
	"time"
)

type VaildPhoneService struct {
	OperationType int `form:"operation_type" json:"operation_type"`
	Phone string `form:"phone" json:"phone"`
	Code  string `form:"code" json:"code"`
}

type GetCodeService struct {
	Phone string `form:"phone" json:"phone"`
}


func (service *VaildPhoneService) Vaild(authorization string) serializer.Response {
	var phone string
	var openid string
	code := e.Success
	phone = service.Phone
	claims,_:= util.ParseToken(authorization)
	openid = claims.OpenID
	if service.OperationType == 1 {
		//1.绑定手机
		if err := cache.RedisClient.Get("code").Err(); err != nil{
			fmt.Println(err)
		}
		RedisCode := fmt.Sprintf("%s",cache.RedisClient.Get("code"))[10:]
		if  RedisCode != service.Code {
			fmt.Println(RedisCode, service.Code)
			code = e.ErrorMsgCode
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		if err := model.DB.Model(model.User{}).Where("open_id=?", openid).Update("phone", phone).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else if service.OperationType == 2 {
		//2.解绑手机
		if err := cache.RedisClient.Get("code").Err(); err != nil{
			fmt.Println(err)
		}
		RedisCode := fmt.Sprintf("%s",cache.RedisClient.Get("code"))
		if  RedisCode != service.Code {
			code = e.ErrorMsgCode
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		if err := model.DB.Model(model.User{}).Where("open_id=?", openid).Update("phone", "" ).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	//获取该用户信息
	var user model.User
	if err := model.DB.First(&user).Where("open_id = ?",openid).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

//发送短信
func (service *GetCodeService) SendMsg() serializer.Response {
	code := e.Success
	rand.Seed(time.Now().UnixNano())
	codeInt := rand.Intn(10000)
	codeString := strconv.Itoa(codeInt)
	if err := cache.RedisClient.Set("code", codeString, 0).Err(); err != nil{
		logging.Info(err)     //将code存入redis中
	}
	if err := cache.RedisClient.Get("code").Err(); err != nil{
		logging.Info(err)    //将code从redis拿出来
	}
	credential := common.NewCredential(  //创建第一个实例对象，登陆用
		conf.TxSecretId,
		conf.TxSecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1" 			//SDK 默认用 TC3-HMAC-SHA256 进行签名，非必要请不要修改该字段
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()  					//创建第二个实例对象，发送信息用
	request.SmsSdkAppid = common.StringPtr(conf.TxSmsSdkAppid)		//短信应用 ID: 在 [短信控制台] 添加应用后生成的实际 SDKAppID
	request.Sign = common.StringPtr(conf.TxSmsSign)  				//短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，可登录 [短信控制台] 查看签名信息 */
	request.TemplateParamSet = common.StringPtrs([]string{codeString})   //放{1}参数，验证码
	request.TemplateID = common.StringPtr(conf.TxTemplateID)
	request.PhoneNumberSet = common.StringPtrs([]string{"+86"+service.Phone})  //发送的号码
	response, err := client.SendSms(request)     			// 通过 client 对象调用想要访问的接口，需要传入请求对象
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		code = e.ErrorSendMsg
		return serializer.Response{
			Status: code,
			Data:   fmt.Sprintf("%s", err),
			Msg:    e.GetMsg(code),
		}
	}
	Info , err := json.Marshal(response.Response)
	InfoStr := fmt.Sprintf("%s", Info)
	if err!=nil {
		code = e.ErrorSendMsg
		return serializer.Response{
			Status: code,
			Data:   err,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   InfoStr,
		Msg:    e.GetMsg(code),
	}
}