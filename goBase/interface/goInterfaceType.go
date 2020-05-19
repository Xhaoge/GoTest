package main

import "fmt"

type People interface {
	run()
	read()
}

type xhaoge struct {
	name  string
	sport string
	book  string
}

func (x *xhaoge) run() {
	fmt.Println("%s 开始跑步...", x.name)
}

func (x *xhaoge) read() {
	fmt.Println("%s 喜欢读%s", x.name, x.book)
}

func findType(xxx interface{}) {
	switch xxx.(type) {
	case string:
		fmt.Println("string")
	case People:
		fmt.Println("people")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	var xx People
	xx = &xhaoge{
		name:  "xhaoge",
		sport: "baseketball",
		book:  "富翁的书....",
	}
	fmt.Println(xx)
	// findType(xx)

	val, ok := xx.(*xhaoge)
	if !ok {
		fmt.Println("not ok")
	}
	fmt.Println(val)
	findType(xx)
}
