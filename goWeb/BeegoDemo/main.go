package main

import (
	_ "goWeb/BeegoDemo/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	beego.Info("beego test")
}
