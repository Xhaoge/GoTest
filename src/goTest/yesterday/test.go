package main

import "fmt"

func addmatch() {
	fmt.Println("nima")
}

type Xhaoge struct {
	name  string
	age   int
	title string
}

func main() {
	fmt.Println("nima!")
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" // 多重赋值

	area = LENGTH * WIDTH
	fmt.Println(area)
	fmt.Println("面积为：%d", area)
	//Println("")
	fmt.Println(a, b, c)
	addmatch()

	var gg Xhaoge
	gg.name = "孙浩"
	gg.age = 23948
	gg.title = "帅气"
	fmt.Println(gg.name)
	printXhaoge(gg)
}

func printXhaoge(g Xhaoge) {
	fmt.Println(g.name)
	fmt.Println(g.age)
	fmt.Println(g.title)
}
