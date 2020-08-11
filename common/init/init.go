package init

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/jinzhu/gorm"
)

var (
	ConfigInit  Yaml
	EtcdClient  *clientv3.Client
	MysqlClient *gorm.DB
)

func InitAll() (err error) {
	if ConfigInit, err = InitConfig(); err != nil {
		return
	}

	if EtcdClient, err = InitEtcd(); err != nil {
		return
	}

	if MysqlClient, err = InitMysql(); err != nil {
		return
	}
	return
}
