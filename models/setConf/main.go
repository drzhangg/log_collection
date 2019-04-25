package main

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"context"
	"time"
)

const (
	EtcdKey = "/oldboy/backend/secskill/product"
)

//定义秒杀商品信息
type SecInfoConf struct {
	ProductId int
	StartTime int
	EndTime   int
	Status    int
	Total     int
	Left      int
}

func SecLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed,err:", err)
		return
	}

	fmt.Println("connect success")
	defer cli.Close()

	var SecInfoConArr []SecInfoConf
	SecInfoConArr = append(
		SecInfoConArr,
		SecInfoConf{
			ProductId: 1028,
			StartTime: 1556363763,
			EndTime:   1556450163,
			Status:    0,
			Total:     1000,
			Left:      1000,
		},
	)
	SecInfoConArr = append(
		SecInfoConArr,
		SecInfoConf{
			ProductId: 1027,
			StartTime: 1556363763,
			EndTime:   1556450163,
			Status:    0,
			Total:     2000,
			Left:      1000,
		},
	)

	data, err := json.Marshal(SecInfoConArr)
	if err != nil {
		fmt.Println("json failed,", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed,err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed,err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	SecLogConfToEtcd()
}
