package main

import (
	"github.com/astaxie/beego"
	"github.com/drzhangg/secKill/models"
)

func main() {
	err := models.InitConfig()
	if err != nil {
		panic(err)
		return
	}
	err = models.InitSec()
	if err != nil {
		panic(err)
		return
	}

	beego.Run()
}
