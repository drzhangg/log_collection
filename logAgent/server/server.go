package server

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"sync"
)

var (
	err        error
	EtcdClient *clientv3.Client
	MutexLock  sync.Mutex
	wg         sync.WaitGroup
)

func init() {
	if err = InitConfig(); err != nil {
		panic(fmt.Sprintf("init server failed, err:%v", err))
	}
}

//初始化配置
func InitConfig() (err error) {
	//初始化配置文件

	//初始化etcd

	//初始化nsq

	//初始化log

	//读取日志搜集信息
	return
}

func Run() {
	//监听etcd变化
	wg.Wait()
}
