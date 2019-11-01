package main

import (
	_ "goMyblog/BeegoDemo/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	beego.Info("beego test")
}
