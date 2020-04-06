package controllers

import (
	"fmt"
	"log"
	"goWeb/myblog/models"
)

type UpdateArticleController struct {
	BaseController
}

type DeleteArticleController struct {
	BaseController
}

func (this *UpdateArticleController) Get(){
	id,_ := this.GetInt("id")
	fmt.Println(id)
	//获取id所对应的文章信息；
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id
	this.TplName = "write_article.html"
}


func (this *UpdateArticleController) Post() {
	id,_ := this.GetInt("id")
	fmt.Println("post id =",id)
	//获取浏览器传输的数据，通过表单的name 属性获取值；
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	//实例化model 修改数据库；
	art := models。Article{id,title,tags,short,content,"",0}
	_,err := models.UpdateArticle(art)
	//返回数据给浏览器；
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code":1,"message":"更新成功；"}
	} else {
		this.Data["json"] = map[string]interface{}{"code",0,"message":"更新失败；"}
	}
}


func (this *DeleteArticleController) Get(){
	artId,_ := this.GetInt("id")
	fmt.Println("delete id = ",artId)
	_,err := models.DeleteArticle(artId)
	if err != err {
		log.Println(err)
	}
	this.Redirect("/",302)

}