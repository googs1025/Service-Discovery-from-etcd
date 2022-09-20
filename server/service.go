package server

import (
	"context"
	"fmt"
	"server/rpc"
)

type Server struct {

}


func Hello(ctx context.Context, request *rpc.HelloRequest) (*rpc.HelloResponse, error) {
	resp := rpc.HelloResponse{
		Hello: "客户端已成功调用服务端的函数！"
	}
	return &resp, nil
}

func Register(ctx context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {

	resp := rpc.RegisterResponse{}
	resp.Uid = fmt.Sprintf("%s.%s", request.GetName(), request.GetPassword())
	return &resp, nil

}
