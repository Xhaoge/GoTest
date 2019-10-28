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
	fmt.Printf("\n%s %s\n", p.firstName, p.lastName)
}

//递归调用函数；
func test(a int){
	if a==1{
		fmt.Println("终止函数")
		return
	}
	test(a-1)
	fmt.Println("a = ",a)
}


func testAdd(i int) int {
	if i==1{
		return 1
	}
	return i+testAdd(i-1)
}

#todo 回调函数 函数的一个参数是函数类型；


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
	// string 操作；
	name := "Hello World"
    for i:= 0; i < len(name); i++ {
        fmt.Printf("%d ", name[i])
	}
	fmt.Printf("\n")
	for i:= 0; i < len(name); i++ {
        fmt.Printf("%c ", name[i])
	}

	test(3)
	z := testAdd(10)
	fmt.Println("zz = ",z)

}
