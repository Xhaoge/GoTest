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
