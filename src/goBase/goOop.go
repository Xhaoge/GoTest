package main
/* 面向对象 test  go并不是一个纯面向对象的编程语言，在go的面向对象，结构体代替了类*/

import "fmt"

type Xhaoge struct{
	name string
	age int
	work string
	feature string
	xbase
}

type xbase struct{
	high int
	des string
}

func (x *Xhaoge) Yutu() {
	fmt.Printf("my name:%s, age:%d, work:%s, in the feature:%s\n",x.name,x.age,x.work,x.feature)
	if x.age >10{
		x.age=9
	}	
}

type haoge struct{
	xhaoges []Xhaoge
}

func (w *haoge)print_haoge() {
	fmt.Println("this is print_haoge")
	for _,x := range w.xhaoges {
		fmt.Printf("struct:%v,  type:%T, name:%s,work:%s,des:%d\n",x,x,x.name,x.work,x.high)
	}
}


func main(){
	fmt.Println("test OOP")
	des := xbase{171,"很帅"}
	xx := Xhaoge{name:"xhaoge",age:13,work:"yuetu",feature:"充满希望",xbase:des}
	xx.Yutu()
	fmt.Printf("my name:%s, age:%d, work:%s, in the feature:%s\n",xx.name,xx.age,xx.work,xx.feature)
	xxx :=Xhaoge{name:"xhao",age:17,work:"yuetu",feature:"财富自由",xbase:des}

	w := haoge{
		xhaoges: []Xhaoge{xx,xxx},
	}
	w.print_haoge()

}
