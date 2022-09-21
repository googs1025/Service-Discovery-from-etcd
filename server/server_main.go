package main

import (
	"context"
	"flag"
	"fmt"
	"main/rpc"

	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"


)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", port)


	stopC := make(chan os.Signal, 1)
	signal.Notify(stopC, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-stopC
		_ = EtcdUnRegister(addr)
		if value, ok := s.(syscall.Signal); ok {
			os.Exit(int(value))
		} else {
			os.Exit(0)
		}

	}()

	err := EtcdRegister(addr)

	if err != nil {
		panic(err)
	}

	connection, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor()))
	rpc.RegisterServerServer(grpcServer, Server{})

	log.Println("service start port %d\n", port)
	if err := grpcServer.Serve(connection); err != nil {
		panic(err)
	}



}

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("call %s\n", info.FullMethod)
		resp, err = handler(ctx, req)
		return resp, err
	}
}