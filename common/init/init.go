package init

import "github.com/coreos/etcd/clientv3"

var (
	ConfigInit Yaml
	EtcdClient *clientv3.Client
)

func InitAll() (err error) {
	if ConfigInit, err = InitConfig(); err != nil {
		return
	}

	if EtcdClient, err = InitEtcd(); err != nil {
		return
	}
	return
}
