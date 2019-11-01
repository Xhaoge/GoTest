package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	beego.Info("第一个beego案例")
	beego.Info("代码修改....")
	fmt.Println("xhaoge")
	beego.Run("localhost:8080")
}
