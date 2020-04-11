package pkg

import (
	"fmt"
	"github.com/Xhaoge/sh/myhttp"
)

func main()  {
	fmt.Println("hello world......")
	pkgNumber := myhttp.GetRandomStr(4)
	fmt.Print("pkgNumber: ",pkgNumber)
}
