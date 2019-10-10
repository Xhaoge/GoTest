package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func1:处理最基本的get
func func1(c *gin.Context) {
	//回复一个2000k，在client 的http-get的resp 的body中获取数据
	c.String(http.StatusOK, "test1 ok")
}

// func2:处理最基本的post
func func2(c *gin.Context) {
	//回复一个200 ok，在client的http-post的body中回去数据
	c.String(http.StatusOK, "test2,nima 怎么才来；")
}

// func3：处理带参数的额path-get
func func3(c *gin.Context) {
	// 回复一个200 ok
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数：%s %s   test3 ok", name, passwd)
}

//func4：处理带出的path-post
func func4(c *gin.Context) {
	// 回复一个200 ok
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数：%s %s  test4 ok", name, passwd)
}

// func5：注意“：” 和“*”的区别
func func5(c *gin.Context) {
	// 回复一个200 ok
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数：%s %s  test5 ok", name, passwd)

}

type Person struct {
	msg   string
	extra string
}

//func6：参数是从body中获得，而不是url；
func func6(c *gin.Context) {

	var person Person
	//c.Writer.Write([]byte("tianja"+person.msg))
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"status": "200",
		"msg":    person.msg,
		"nick":   nick,
		"extra":  person.extra,
	})
}

func main() {
	fmt.Println("nima")
	//注册一个最基本的路由器
	router := gin.Default()
	//最基本的用法
	router.GET("/test1", func1)
	router.POST("/test2", func2)
	// 注意："："必须要匹配，"*"选择匹配，即存在就匹配，不存在可以不用考虑
	router.GET("/test3/:name/:passwd", func3)
	router.POST("/test4/:name/:passwd", func4)
	router.GET("/test5/:name/*passwd", func5)
	//所有参数是需要从body中获得
	router.POST("/test6", func6)
	// 绑定端口是8888
	router.Run(":8888")
}
