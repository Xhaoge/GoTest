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
	beego.Router("/article/add",&controllers.AddArticleController{})
}
