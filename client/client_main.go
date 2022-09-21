package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"golanglearning/new_project/ServiceDiscoveryByETCD/client/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	etcdURL = "http://localhost:2379"
	serviceName = "example/server"

)

func main() {

	etcdClient, err := clientv3.NewFromURL(etcdURL)
	if err != nil {
		panic(err)
	}
	etcdResolver, err := resolver.NewBuilder(etcdClient)
	connection, err := grpc.Dial(
		fmt.Sprintf("etcd:///%s", serviceName),
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
	)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	// 主要调用server端的Hello方法
	ServerClient := rpc.NewServerClient(connection)
	for {
		helloRespone, err := ServerClient.Hello(context.Background(), &rpc.HelloRequest{})
		if err != nil {
			fmt.Printf("err: %v", err)
			return
		}

		log.Println(helloRespone, err)
		time.Sleep(500 * time.Millisecond)
	}


}
