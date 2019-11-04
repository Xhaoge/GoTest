package models

import "goWeb/myblog/utils"

type Article struct{
	Id	int
	Title string
	Tags string
	Short string
	Content string
	Author string
	Createtime int64
}

//--------------------数据处理-----------
func AddArticle(article Article) (int64,error) {
	i,err := insertArticle(article)
	return i,err
}

//----------------------数据库操作--------------
//插入一篇文章
func insertArticle(Article Article)(int64,error){
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title,article.Tags,article.Short,article.Content,article.Author,article.Createtime)
}

// 根据页码查询文章
func FindArticleWithPage(page int)([]Article,error){
	//从配置文件中共获取每页的文章数量
	num,_ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("-------------->page",page)
}