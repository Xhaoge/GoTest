package main

import "fmt"
/*定义函数返回俩个值得和；*/

func add(num1,num2 int) int {
	var result int
	result = num1+num2
	return result
}

func main(){
	const(
		a=3
		b=40
	)
	var ret int
	if (b<a){
		fmt.Println("a-b:",a-b)
	}else{
		ret = add(a,b)
	}
	fmt.Println(ret)
	fmt.Println("你在哪里")
	fmt.Println("我不知道你在说什么 哦   还是那个地点  那条街哦 ；")
	fmt.Println("糖糖  你在搞啥子哦 你滚哦")
}