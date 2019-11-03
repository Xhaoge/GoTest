package controllers

import (
	"fmt"
	"goWeb/myblog/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post(){
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:",username,"password:",password)

	id := models.QueryUserWithParam(username,password)
	fmt.Println("log 时登陆的账号id为：",id)
	if id > 0{
		this.Data["json"] = map[string]interface{}{"code":1,"message":"登陆成功"}
	}else{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"登陆失败"}
	}
	this.ServeJSON()
}


