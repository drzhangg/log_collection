package models

import (
	"fmt"
	"github.com/astaxie/beego"
)

var (
	secSkillConf = &SecSkillConf{}
)

type RedisConf struct {
	redisAddr        string
	redisMaxIdle     int
	redisMaxActive   int
	redisIdleTimeout int
}

type EtcdConf struct {
	etcdAddr string
	timeout  int
}

type SecSkillConf struct {
	redisConfig RedisConf
	etcdConf    EtcdConf
	logPath     string
	logLevel    string
}

func InitConfig() (err error) {

	beego.LoadAppConfig("ini", "conf/app.conf")
	//beego.BConfig.Listen.HTTPAddr = "192.168.0.104"

	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")
	//logs.Debug("read config succ,redis addr:%v", redisAddr)
	//logs.Debug("read config succ,etcd addr:%v", etcdAddr)

	secSkillConf.redisConfig.redisAddr = redisAddr
	secSkillConf.etcdConf.etcdAddr = etcdAddr

	if len(redisAddr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed,redis[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return
	}

	redisMaxIdle, err := beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		err = fmt.Errorf("init config failed,read redis_max_idle error:%v", redisMaxIdle)
		return
	}

	redisMaxActive, err := beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		err = fmt.Errorf("init config failed,read redis_max_active error:%v", redisMaxActive)
		return
	}

	redisIdleTimeout, err := beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		err = fmt.Errorf("init config failed,read redis_idle_timeout error:%v", redisIdleTimeout)
		return
	}

	secSkillConf.redisConfig.redisMaxIdle = redisMaxIdle
	secSkillConf.redisConfig.redisMaxActive = redisMaxActive
	secSkillConf.redisConfig.redisIdleTimeout = redisIdleTimeout

	etcdTimeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed,read etcd_timeout error:%v", etcdTimeout)
	}

	secSkillConf.etcdConf.timeout = etcdTimeout

	secSkillConf.logPath = beego.AppConfig.String("log_path")
	secSkillConf.logLevel = beego.AppConfig.String("log_level")
	return
}
