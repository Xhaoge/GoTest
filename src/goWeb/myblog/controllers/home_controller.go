package controllers

import (
	"fmt"
	"goWeb/myblog/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

type HomeController struct {
	//beego.Controller
	BaseController
}

func (this *HomeController) Get() {
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	page,_ := this.GetInt("page")
	if page <= 0{
		page = 1
	}
	var artList []models.Article
	artList, _ = models.FindArticleWithPage(page)
	this.Data["PageCode"] = 1
	this.Data["HasFooter"] = true
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList,this.IsLogin)
	this.TplName = "home.html"
}

// 判断是否登录
func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginuser")
	fmt.Println("loginuser---->", loginuser)
	if loginuser != nil {
		this.IsLogin = true
		this.Loginuser = loginuser
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}
