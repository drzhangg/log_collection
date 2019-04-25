package main

import (
	"secKill/models"
	_ "secKill/routers"
	"github.com/astaxie/beego"

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

