package controllers

import (
	"fmt"
	"goWeb/myblog/models"
	"time"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

// 处理注册
func (this *RegisterController) Post() {
	// 获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)

	// 注册之前先判断该用户是否已经被注册，如果已经注册，返回错误；
	id := models.QueryUserWithUsername(username)
	fmt.Println("判断时该用户是否已经注册，id = ", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在！amazed"}
		this.ServeJSON()
		return
	}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里的密码 进行判断
	//todo password = utils.MD5(password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	fmt.Println("注册的用户信息 = ", user)
	_, err := models.InsertUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	this.ServeJSON()

}
