package init

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

func InitEtcd() (etcdClient *clientv3.Client, err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   ConfigInit.Etcd.Addr,
		DialTimeout: time.Duration(ConfigInit.Etcd.EtcdDailTimeout) * time.Second,
	})
	if err != nil {
		err = fmt.Errorf("connect etcd failed, err:", err)
		log.Println(err)
		return
	}
	etcdClient = cli
	log.Println("etcd init success")
	return
}
