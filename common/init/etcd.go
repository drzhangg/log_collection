package init

import "github.com/coreos/etcd/clientv3"

func InitEtcd() (etcdClient *clientv3.Client, err error) {
	clientv3.New(clientv3.Config{
		Endpoints:   nil,
		DialTimeout: 0,
	})
	return
}
