package main

import (
	"fmt"
	"github.com/xiaohaogge/sh/myhttp"
)

func main()  {
	fmt.Println("hello world......")
	pkgNumber := myhttp.GetRandomStr(4)
	fmt.Print("pkgNumber: ",pkgNumber)
}
