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

type HomeBlockParam struct{
	Id 			int
	Title 		string
	Tags		[]TagLink
	Short 		string
	Content		string
	Author		string
	CreateTime	string
	//查看文章的地址；
	Link 		string

	// 修改文章的地址；
	Update 		string
	DeleteLink 	string

	//记录是否登陆；
	IsLogin 	string
}

//标签连接；
type TagLink struct{
	TagName		string
	TagUrl		string
}

type HomeFooterPageCode struct {
	HasPre  	bool
	HasNext 	bool
	ShowPage	string
	PreLink		string
	NextLink	string
}


func (this *HomeController) Get() {
	tag := this.GetString("tag")
	fmt.Println("tags:",tag)
	page,_ := this.GetInt("page")
	if page <= 0{
		page = 1
	}
	var artList []models.Article
	if len(tag) > 0 {
		// 按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		this.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		// 设置分页
		artList, _ = models.FindArticleWithPage(page)
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasNext"] = true
	}

	// artList, _ = models.FindArticleWithPage(page)
	// this.Data["PageCode"] = 1
	// this.Data["HasFooter"] = true
	fmt.Println(artList)
	
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList,this.IsLogin)

	this.TplName = "home.html"
}

// 判断是否登录
// func (this *BaseController) Prepare() {
// 	loginuser := this.GetSession("loginuser")
// 	fmt.Println("loginuser---->", loginuser)
// 	if loginuser != nil {
// 		this.IsLogin = true
// 		this.Loginuser = loginuser
// 	} else {
// 		this.IsLogin = false
// 	}
// 	this.Data["IsLogin"] = this.IsLogin
// }


// // -------------翻页---------
// // page 是当前的页数；
// func ConfigHomeFooterPageCode (page int) HomeFooterPageCode {
// 	pageCode := HomeFooterPageCode{}
// 	// 查询出总的条数；
// 	num := GetArticleRowsNum()
// 	// 从配置文件中共读取每页显示的条数
// 	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
// 	// 计算出总的页数；
// 	allPageNum := (num-1)/pageRow +1
// 	pageCode.ShowPage = fmt.Sprintf("%d/%d",page,allPageNum)
// 	// 当前页数小于等于1，那么上一页的按钮不能点击；
// 	if page <= 1{
// 		pageCode.HasPre = false
// 	}else {
// 		pageCode.HasPre = true
// 	}
// 	// 当前页数 大于等于总页数，那么下一页的按钮不能点击；
// 	if page  >= allPageNum {
// 		pageCode.HasNext = false
// 	} else {
// 		pageCode.HasNext = true
// 	}
// 	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
// 	pageCode.NextLink = "?page=" + strconv.Itoa(page+1)
// 	return pageCode
// }