package main

import (
	_ "goWeb/myblog/routers"
	"goWeb/myblog/utils"

	"github.com/astaxie/beego"
)

func main() {
	beego.Info("this is my blog")
	utils.InitMysql()
	beego.Run()
}
