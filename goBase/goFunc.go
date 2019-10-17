package main

import "fmt"

func maxNum(a,b int) int {
	var result int
	if a>b{
		result=a
	}else{
		result=b
	}
	fmt.Printf("最大数为：%d\n",result)
	return result

}


func main(){
	fmt.Println("这个是函数test")
	maxNum(3,6)
}