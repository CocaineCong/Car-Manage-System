# Car-Manage-System 车辆信息管理平台

#### 此项目为前后端分离项目

#### 前台展示页面是用`小程序`搭建，后台展示页面使用`vue`搭建。 小程序前端代码地址：[小程序端](https://github.com/Sakura-MI/Car-system)

#### 后端主部分使用`Golang`的`Gin`框架，还有一部分是封装了`OCR`识别`Python`的`Flask`框架。 Python端代码[python](https://github.com/Poootato/SoftwareProject)

## 前言

本人目前大三，这是我们的软工实践大作业。喜欢的话，别忘记左上角`star`一下

## 项目依赖

golang部分：

- Gin
- Gorm
- mysql
- redis
- mongodb
- ini
- jwt-go
- websocket

python部分：

- flask
- opencv-python
- pytorch

使用的SDK

- 七牛云存储
- 腾讯云短信

## 开发环境


后 端：Golang v1.15、Python v3.7

前 端：微信小程序基础库 v2.16.0、Vue v3.3.0

算 法 ： Pytorch v1.7.1、Cuda v11.0 

数 据 库 ： MySQL v5.7.30、MongoDB v4.4.6、Redis v4.0.9

短信服务 ：腾讯云短信

文件存储 ：七牛云存储

服 务 器 ： 阿里云服务器

## Go目录结构

```
CarDemo1/
├── api
├── conf
├── middleware
├── model
├── pkg
│  ├── e
│  ├── logging
│  ├── util
├── routes
├── serializer
├── servive
│  ├── ws
│  	   ├── e
│      ├── model
└── upload

```

- api：用于定义接口函数
- conf：用于存储配置文件
- middleware：应用中间件
- model：应用数据库模型
- pkg / e：封装错误码
- pkg / logging: 日志打印
- pkg / util：工具函数
- serializer：将数据序列化为 json 的函数
- routes ：路由逻辑处理
- service：接口函数的实现
- service/ws：聊天功能的实现
- service/ws/e：聊天功能的所需要的状态码
- service/ws/model：聊天功能的模型



## Go所需依赖准备：

项目在启动的时候依赖以下环境变量，可以在项目conf目录下创建config.ini 文件设置环境变量便于使用

```
#debug开发模式,release生产模式
[service]
AppMode = debug
HttpPort = :3000 
# 运行端口号 3000端口
[mysql]
Db = mysql
DbHost = "" 
# mysql的ip地址
DbPort = ""
# mysql的端口号,默认3306
DbUser = ""
# mysql user
DbPassWord = ""
# mysql password
DbName = ""
# 数据库名字

[redis]
RedisDb = ""
# redis 名字
RedisAddr = ""
# redis 地址
RedisPw = ""
# redis 密码
RedisDbName = ""
# redis 数据库名

[wechat] 
# 这两个是在微信开发者后台拿的
APPID = ""
SECRET = ""

[txsms]
# 这些是腾讯云短信的服务,具体可查文档
SecretId = ""
SecretKey = ""
TxSmsSign = ""
# 短信签名
TxSmsSdkAppid = ""
TxTemplateID = ""
# 短信模板ID

[qiniu]
# 七牛云的存储，具体也可查文档
AccessKey = ""
SerectKey = ""
Bucket = ""
QiniuServer = ""

[MongoDB]
MongoDBName =  ""
MongoDBAddr = ""
MongoDBPwd = ""
MongoDBPort = ""
```



## 简要说明

1. mysql 数据库一定要有，不然跑不起来。
2. redis 是用在存储腾讯云短信的验证码，具体有效期限可以直接设置，腾讯云只是用在短信发送，没有对运行没有很大的关系。
3. mongodb 使用在聊天模块的信息存储，可以上拉下拉聊天记录。
4. 本项目的图片我都是存放到七牛云上的，有必要了解一下七牛云的存储。（百度或是看文档就可以看懂了的）



## Golang端运行

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

下载依赖

```
go mod tidy
```

直接运行项目就可以了

项目运行后启动在 3000 端口



## Python目录结构

```
flask_ocr/
├── checkpoints
├── detect
└── recognize
```



- checkpoints ：存放模型文件
- detect ：CTNP 网络
- recognize ： RCNN网络
- app.py ：接口文件

## Python端运行

pip install 完所需要的第三方库就可以运行了，但是要特别注意一下Pytorch和Cuda的版本！



## 项目部分功能说明与展示

## 1. 实现目标：

1.	用户可通过拍照识别车牌号进行绑定车牌号，也能通过车牌号找到对应的车主。
2.	车牌冲突可进行申诉反馈。
3.	可通过文本消息提醒、在线聊天、短信提醒等形式与对方车主产生联系。
4.	强大的社区模块，支持闲置物品的交互买卖。
5.	实时获取充电桩信息，方便用户选择。
6.	除基本管理操作，后台还可对车流量进行实时监控。

## 2. 功能介绍

***

### 2.1 主体部分




<img src="https://img-blog.csdnimg.cn/3fec8fb4770441f5924d7bfcbfed9b5d.png" width="45%" >


### 2.2 用户模块

个人中心是可以对用户个人的信息进行修改、由于是用微信登陆，所以姓名和头像是读取微信的头像和名字。所以名字和头像是不支持修改的。

但是手机号、邮箱号、车辆是可以进行解绑定的。
用户可以通过绑定自己的车牌号来管理自己的车辆。

车牌，我们提供了一个ocr的算法接口，可以对车牌进行识别，然后返回车牌信息进行绑定车辆。

- 个人信息
- 绑定邮箱
- 绑定手机
- 绑定车牌

***

<img src="https://img-blog.csdnimg.cn/20210601163755462.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70">
<img src="https://img-blog.csdnimg.cn/20210601165420913.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70">


<img src="https://img-blog.csdnimg.cn/20210601164723199.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70" width="47%">
<img src="https://img-blog.csdnimg.cn/120deb384ba54fb8b77564ed0fd27a33.png" width="47%">




### 2.3 社区模块

推荐模块、亲友圈、闲来康康、我的世界等。

- 我的世界模块可以查看到用户个人发布的帖子。
- 帖子详情、可以对帖子进行评论、点赞等操作。
- 帖子发送，用户可以通过话题进行发布帖子。

***

<img src="https://img-blog.csdnimg.cn/20210601165725598.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70">
<img src="https://img-blog.csdnimg.cn/20210601165748755.png?x-oss-
process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70">

<img src="https://img-blog.csdnimg.cn/2021060200261436.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70">
<img src="https://img-blog.csdnimg.cn/20210601165907562.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70">


### 2.4 聊天模块以及充电功能

- 聊天功能，实现实时聊天。
- 系统消息，系统可有针对性的对其进行发送信息。
- 用户反馈，可以进行评论举报、聊天举报、车牌申诉等功能。
- 充电桩查询，我们用爬虫将学校充电桩的情况进行爬取，使得用户能够查看充电桩的情况。

***

<img src="https://img-blog.csdnimg.cn/bc5c6bda393c413290799a66fb647e2f.png">

<img src="https://img-blog.csdnimg.cn/92764760c9e649eca34cb11d2676cbcf.png">



### 2.5 算法方面

算法部分的结果都是通过flask框架进行api接口的返回。

***

#### 2.5.1 FasterRCNN网络车牌识别

<img src="https://img-blog.csdnimg.cn/20210602004440282.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70" width="47%">
<img src="https://img-blog.csdnimg.cn/20210602004543510.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70" width="47%">

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210602004622894.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)


#### 2.5.2 YOLOV5 车辆识别

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210602005316890.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20210602005411311.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)

### 2.6 后台管理模块

后台模块相对简单，并没有设计到比较多的功能，后需再进行完善。

   - 用户模块管理
     - 车辆模块管理
     - 反馈信息管理
     - 车流监控管理

***

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601173016305.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)
可对用户进行拉黑、封号处理

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601172935299.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)
可下架、修改用户的帖子信息。


![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601172938842.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)
可对用户的车辆进行处理、更换车牌号等



## 3. 总结

- go的ws也有涉及。
- gorm的多对多也有了深入的了解。还有后端的一些逻辑结构。
- 熟悉了腾讯云短信，七牛云存储，阿里云服务器的一些操作。

这一次的算法方面

- FasterRCNN的车牌识别
- YOLO网络的车辆检测



喜欢的小伙伴可以右上角一个`Star`噢~

