package main

import (
	"fmt"
	"github.com/Xhaoge/sh/myhttp"
	"goUtils"
)

func main()  {
	fmt.Println("hello world......")
	pkgNumber := myhttp.GetRandomStr(4)
	fmt.Print("pkgNumber: ",pkgNumber)
	xx := goUtils.Xhao{Name:"xxxx"}
	fmt.Println(xx)
}
