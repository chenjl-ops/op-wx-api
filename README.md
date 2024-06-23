# op-wx-api
---
## Description
go实现微信公众号相关

- 微信公众号开发者接口实现
- 基于公众号实现消息发送,图片等
- 目前实现消息自动回复,收到什么回复什么

## 配置
```
设置环境变量:
RUNTIME_ENV=test 
RUNTIME_APP_NAME=op-wx-api
RUNTIME_CONFIG_URL=http://nacosconf.url

Nacos Conf:

{
    "SERVER_RUN_PORT": 80, #服务监听端口
    "WX_TOKEN":"xxxx", #微信管理后台配置
    "WX_APP_ID": "xxxxx", #微信管理后台获取
    "WX_APP_SECRET": "xxxx", #微信管理后台配置获取
    "WX_ENCODEING_AES_KEY": "xxx" #微信加解密消息配置
}

```

## 启动
```
由于微信公众号开发者回调只支持80端口,所以启动服务监听80

go mod tidy
go run ./cmd/app/main.go

```

## build
```
docker build harbor.url/op-wx-api:v1 ./
docker push harbor.url/op-wx-api:v1

```
## docker run
```
docker run -d --name op-wx-api --network host --restart on-failure -e RUNTIME_ENV=test -e RUNTIME_APP_NAME=op-wx-api -e RUNTIME_CONFIG_URL=http://10.2.0.167 harbor-public.123go.club/public-test/go-sso-api:v4

```
