package main

import "fmt"

/*这是一个关于变量的test */

func main(){
	var age int=3
	var nima bool = true;
	var a string="我要怎么才能不爱你。。"
	var b string="滚"
	c:="我尼玛"
	d :=4
	
	fmt.Println(a,b,c,d,nima,age,"niam"+"我该怎么潇洒的活在成都；；；；")
	
	const length int=5
	const width=8
	var area int 
	const e,f,g =1,false,"str"
	area = length*width
	fmt.Println("面积为：%d",area)

	fmt.Println(e,f,g)
	
	const(
		unknown ="qwrqe"
		female =len(unknown)
		Male = 2
	)
	fmt.Println(female,unknown)
	
	
	var firs int = 100
	var seco int = 37
	firs++
	fmt.Println(firs+seco,firs-seco,firs/seco,firs%seco,firs)
	
}