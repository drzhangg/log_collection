package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	etcd_client "go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

var (
	redisPool  *redis.Pool
	etcdClient *etcd_client.Client
)

func initRedis() (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     secSkillConf.redisConfig.redisMaxIdle,
		MaxActive:   secSkillConf.redisConfig.redisMaxActive,
		IdleTimeout: time.Duration(secSkillConf.redisConfig.redisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secSkillConf.redisConfig.redisAddr)
		},
	}

	conn := redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed,err:%v", err)
		return
	}

	return
}

func initEtcd() (err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{secSkillConf.etcdConf.etcdAddr},
		DialTimeout: time.Duration(secSkillConf.etcdConf.timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed,err:", err)
		return
	}

	etcdClient = cli
	return
}

func convertLogLevel(level string) int {
	switch (level) {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return logs.LevelDebug
}

func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = secSkillConf.logPath
	config["level"] = convertLogLevel(secSkillConf.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

//加载sec配置
func loadSecConf() (err error) {
	key := fmt.Sprintf("%s/product", secSkillConf.etcdConf.etcdSecKey)
	resp, err := etcdClient.Get(context.Background(), key)
	if err != nil {
		logs.Error("get [%v] from etcd failed, err:%v", key, err)
		return
	}

	var secProductInfo []SecInfoConf
	for k, v := range resp.Kvs {
		logs.Error("key[%v] value[%v]", k, v)
		err = json.Unmarshal(v.Value,&secProductInfo)
		if err != nil {
			logs.Error("Unmarshal sec product info failed,err :%v",err)
			return
		}
		logs.Debug("sec info conf is [%v]",secProductInfo)
	}

	return
}

//初始化sec
func InitSec() (err error) {
	err = initLogger()
	if err != nil {
		logs.Error("init logger failed,err:%v", err)
		return
	}

	/*
	err = initRedis()
	if err != nil {
		logs.Error("init redis failed,err:%v", err)
		return
	}
	*/

	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed,err:%v", err)
		return
	}

	err = loadSecConf()
	if err != nil {
		return
	}

	logs.Info("init sec success")
	return
}
