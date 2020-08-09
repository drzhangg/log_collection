package init

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Yaml struct {
	Mysql `yaml:"mysql"`
	Etcd  `yaml:"etcd"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Etcd struct {
	Addr            []string `yaml:"addr"`
	ConfigKey       string   `yaml:"config_key"`
	EtcdDailTimeout int      `yaml:"etcd_dail_timeout"`
	EtcdPutTimeout  int      `yaml:"etcd_put_timeout"`
	EtcdGetTimeout  int      `yaml:"etcd_get_timeout"`
}

func InitConfig() (Yaml, error) {
	var yml Yaml
	file, err := ioutil.ReadFile("../conf/common.yml")
	if err != nil {
		log.Panic("readFile failed, err:", err)
	}

	if err := yaml.Unmarshal(file, &yml); err != nil {
		log.Panic("yaml.Unmarshal failed, err:", err)
	}
	return yml, err
}
