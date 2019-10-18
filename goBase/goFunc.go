package main

import (
	"fmt"
)

func maxNum(a, b int) int {
	var result int
	if a > b {
		result = a
	} else {
		result = b
	}
	fmt.Printf("最大数为：%d\n", result)
	return result
}

//defer 函数的适用；在离开所有方法时执行defer，报错的时候也会执行；
/*func ReadWrite() bool {
	file.Open("../files/test.txt")
	defer file.Close()
	if failureX {
		return false
	}
	if failureY {
		return false
	}
	return true
}*/

func nameDefer(n string) {
	name := n
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}

type person struct {
	firstName string
	lastName  string
}

func (p person) getName() {
	fmt.Printf("%s %s\n", p.firstName, p.lastName)
}

func main() {
	defer fmt.Println("这是最后才执行的一个defer 吗？")
	fmt.Println("这个是函数test")
	maxNum(3, 6)
	zz := person{
		firstName: "xhaoge",
		lastName:  "jiayou",
	}
	defer zz.getName()
	fmt.Println("welcome,")

}
