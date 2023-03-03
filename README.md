# MiniDouyin
使用gin＋gorm+redis+mysql实现一个极简版抖音（字节青训营项目）

## 项目简介
`使用Go语言实现了一个极简版抖音，实现了抖音的三大模块：基础模块、互动模块、社交模块。`

`1、基础模块实现了视频的发布以及获取feed流接口、用户的注册登录以及鉴权接口、用户信息获取接口、获取发布列表接口；`

`2、互动模块实现了视频的点赞操作接口、评论操作接口、获取喜欢视频列表以及视频评论列表接口：`

`3、社交模块实现了用户间的关注操作接口、获取关注列表、粉丝列表、好友列表接口、发送消息接口以及获取聊天记录接口；`

## 项目要点
`1、Redis提供缓存`

`2、使用JWT作web服务拦截器，实现用户鉴权`

`3、视频采用OSS云存储，加快访问速度的同时缓解数据库压力`

`4、评论聊天内容进行铭感词过滤`

`5、用户隐私信息采用加密处理，可大幅度降低数据库泄露的危害`

`6、使用goSocket实现用户聊天`

## 数据库结构设计

![数据库](./assets/Diagram%201.jpg)

## 项目结构

```
│  go.mod
│  go.sum
│  main.go
│  README.md
│
├─.idea
│      .gitignore
│      MiniDouyin.iml
│      modules.xml
│      vcs.xml
│      workspace.xml
│
├─api
│      comment.go
│      favorite.go
│      message.go
│      relation.go
│      user.go
│      video.go
│
├─assets
│      Diagram 1.jpg
│
├─config
│      config.go
│
├─db
│      dbConnect.go
│      dbConnect_test.go
│      redisConnect.go
│      redisConnect_test.go
│
├─middleware
│      JWT.go
│
├─model
│      comment.go
│      common.go
│      favorite.go
│      favorite_test.go
│      message.go
│      relation.go
│      user.go
│      video.go
│
├─router
│      router.go
│
├─rpc
│  └─rpcpb
└─utils
        ffmpeg.go
        md5.go
        qiNiuCloud.go
        qiNiuCloud_test.go
        resolveError.go
        statusMsg.go

```

## 项目部署

`1、下载并安装依赖`
```gitignore
git clone git@github.com:ShuXCoding/MiniDouyin.git
```
```go
sudo go mod tidy
```
`2、配置config.go文件`

`3、运行项目`
```go
sudo go run main.go
```
