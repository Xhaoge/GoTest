package routers

import (
	"goWeb/myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/exit",&controllers.ExitController{})
	//写文章
	beego.Router("/article/add",&controllers.AddArticleController{})
	//显示文章内容
	beego.Router("/article/:id",&controllers.ShowArticleController{})
	//更新文章内容
	beego.Router("/article/update",&controllers.UpdateArticleController{})
	// 标签
	beego.Router("/tags",&controllers.TagsController{})
	//相册
	beego.Router("/album",&controllers.AlbumController{})
	//文件上传
	beego.Router("/upload",&controllers.UploadController{})
	//关于我
	beego.Router("/aboutme",&controllers.AboutMeController{})
}
