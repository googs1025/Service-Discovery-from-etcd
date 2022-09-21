package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
	"main/api"
)



// etcd客户端
var etcdClient *clientv3.Client

// EtcdRegister 注册实例
func EtcdRegister(addr string) error {

	log.Printf("etcdRegister %s\b", addr)
	// etcd客户端
	etcdClient, err := clientv3.NewFromURL(etcdUrl)
	if err != nil {
		fmt.Println(api.EtcdConnectionError)
		return err
	}

	em, err := endpoints.NewManager(etcdClient, serviceName)
	if err != nil {
		fmt.Println(api.EtcdConnectionError)
		return err
	}

	ctx := context.Background()

	// 需要为实例续期
	lease, err := etcdClient.Grant(ctx, ttl)
	if err != nil {
		fmt.Println(api.EtcdLeaseError)
		return err
	}

	// 加入
	err = em.AddEndpoint(ctx, fmt.Sprintf("%s/%s", serviceName, addr), endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID))
	if err != nil {
		fmt.Println(api.EtcdAddOrDeleteEndpointError)
		return err
	}
	// 需要发送消息，定时续期，查看健康状态
	alive, err := etcdClient.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		fmt.Println(api.EtcdKeepAliveError)
		return err
	}

	// 启一个goroutine查看健康状态。
	go func() {
		for {
			<-alive
			fmt.Println("etcd server keep alive，健康状态检查。")
		}
	}()

	return nil



}

// EtcdUnRegister 解绑实例
func EtcdUnRegister(addr string) error {

	log.Printf("etcdUnRegister %s\b", addr)
	if etcdClient != nil {
		em, err := endpoints.NewManager(etcdClient, serviceName)
		if err != nil {
			fmt.Println(api.EtcdConnectionError)
			return err
		}
		ctx := context.Background()
		// 移除节点
		err = em.DeleteEndpoint(ctx, fmt.Sprintf("%s/%s", serviceName, addr))
		if err != nil {
			fmt.Println(api.EtcdKeepAliveError)
			return err
		}



	}
	return nil

}