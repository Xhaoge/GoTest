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

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:", username, "password:", password)

	id := models.QueryUserWithParam(username, password)
	fmt.Println("log 时登陆的账号id为：", id)
	if id > 0 {
		/*
		因为设置了session后将数据处理设置到cookie，然后到浏览器进行网络请求的的是偶自动带上cookie，
		因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置；
		*/
		this.SetSession("loginuser",username)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登陆成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登陆失败"}
	}
	this.ServeJSON()
}
	}
	this.ServeJSON()
}

