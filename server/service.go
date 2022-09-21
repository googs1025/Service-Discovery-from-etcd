package main

import (
	"context"
	"fmt"
	"main/rpc"
	"strconv"
)

// TODO: 测试用例没有写！

type Server struct {

}


func (s Server) Hello(ctx context.Context, request *rpc.HelloRequest) (*rpc.HelloResponse, error) {
	resp := rpc.HelloResponse{
		Hello: "客户端已成功调用服务端的函数！",
		Port: strconv.Itoa(Port),
	}
	return &resp, nil
}

func (s Server) Register(ctx context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {

	resp := rpc.RegisterResponse{}
	resp.Uid = fmt.Sprintf("%s.%s", request.GetName(), request.GetPassword())
	return &resp, nil

}
