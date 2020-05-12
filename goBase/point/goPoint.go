package main

import "fmt"

type person interface {
	say()
}

type man struct {
	name string
}

func (m man)say(){
	fmt.Println(m.name,"say hello...")
}

func main()  {
	fmt.Println("this is point test")
	var xx man
	xx.name = "xhaoge"

	var _ person = (*man)(nil)
}
