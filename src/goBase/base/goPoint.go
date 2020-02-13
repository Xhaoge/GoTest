package main

import (
	"fmt"
)


func main(){
	fmt.Println("test for point")
	var m int64
	m = 100
	n := &m
	n1 := *n +1

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(*n)
	fmt.Println(n1)
}