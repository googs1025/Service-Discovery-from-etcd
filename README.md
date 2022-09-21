# Service-Discovery-from-etcd
## 使用ETCD+gRPC 实现简单的服务发现与负载均衡功能。

### server端：

EtcdRegister方法：注册实例地址到etcd中，并采用定时续期的方式实现实例的健康检查机制。

EtcdUnRegister方法：从etcd中解绑实例。

### client端：

主要调用server端的方法。


# 如何启服务
## 1.第一步
###需要先用docker 启动etcd实例，并监听2379端口。
## 2.进入server目录，启动server端，并同时启好几个shell且监听不同端口


`cd ./server`

`go run . -port 8081`

`go run . -port 8082`

`go run . -port 8083`

## 3.进入client目录，启动client端

`cd ./client`

`go run . client_main.go`