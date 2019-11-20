package controllers

import (
	"fmt"
	"goWeb/myblog/models"
	"time"
)

type ExitController struct {
	BaseController
}

func (this *ExitController) Get() {
	// 清除该用户登陆状态的数据；
	this.DelSession("loginuser")
	this.Redirect("/", 302)
}

type AddArticleController struct {
	BaseController
}

/*
当访问/add 路径的时候触发AddArticleController的Get方法，
相应的页面时通过TpName
*/

func (this *AddArticleController) Get() {
	this.TplName = "write_article.html"
}

// 通过this.serverjson 去返回字符串；
func (this *AddArticleController) Post() {
	// 获取浏览器传输的数据，通过表单的name属性获取值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("shorts")
	content := this.GetString("content")
	author := this.GetString("author")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	// 实例化model，将他导入到数据库中；
	art := models.Article{0, title, tags, short, content, author, time.Now().Unix()}
	_, err := models.AddArticle(art)

	// 返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		// 无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	this.Data["json"] = response
	this.ServeJSON()

}
