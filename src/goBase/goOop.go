package main
/* 面向对象 test  go并不是一个纯面向对象的编程语言，在go的面向对象，结构体代替了类*/

import "fmt"

type Xhaoge struct{
	name string
	age int
	work string
	feature string
}

func (x *Xhaoge) Yutu() {
	fmt.Printf("my name:%s, age:%d, work:%s, in the feature:%s\n",x.name,x.age,x.work,x.feature)
	if x.age >10{
		x.age=9
	}	
}


func main(){
	fmt.Println("test OOP")

	var xx Xhaoge
	xx = Xhaoge{name:"xhaoge",age:13,work:"yuetu",feature:"充满希望"}
	xx.Yutu()
	fmt.Printf("my name:%s, age:%d, work:%s, in the feature:%s\n",xx.name,xx.age,xx.work,xx.feature)
}
