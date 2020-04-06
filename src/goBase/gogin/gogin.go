package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("test gin")
	// route := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	fmt.Println("r engin ", r)
	r.GET("ping", pa)
	r.Run(":80")
}

func pa(c *gin.Context) {
	c.JSON(200, "hello")
}
