package main

import (
	"fmt"
	"sh/myhttp"
)

func main()  {
	fmt.Println("hello world......")
	pkgNumber := myhttp.GetRandomStr(4)
	fmt.Print("pkgNumber: ",pkgNumber)
}
