package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	// 获取etcd连接
	conn, err := GetConn()
	if err != nil {
		return
	}
	defer conn.Close()
	//put数据
	if err := putData("jack", "jacktaoyang", conn); err != nil {
		fmt.Printf("putData failed,err:%v/n", err)
		return
	}
	getData("jack", conn)
}

func GetConn() (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return nil, err
	}
	return client, nil
}

// putData putData
func putData(key, value string, cli *clientv3.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, key, value)
	defer cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return err
	}
	return nil
}

// getData getData
func getData(key string, cli *clientv3.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	defer cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return err
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	return nil
}
