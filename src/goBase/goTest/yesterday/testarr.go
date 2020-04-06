package main

import (
	"fmt"
	"strconv"
	) 

/* 这是一个关于arr 数组的练习 
	以及 数据int的转换
	string 的转换；
*/

func main(){
	fmt.Println("nima,这是啥数组哦；")
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance[3])
	var n[10]int   //先声明这是一个长度为10 的数组；
	var k,j int
	for k=0;k<10;k++{
		n[k] = k+100
	}

	for j=0;j<len(n);j++{
		fmt.Println(n[j])
	}
	fmt.Println(n)


	var i int32 =12
	var n1 float32 = float32(i)
	var n2 int8 =int8(i)
	var n3 int64 = int64(i)
	//Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	fmt.Println(&i,n1,n2,n3)

	var n4 int8
	var n5 int32
	n4 = int8(i)+ 1
	n5 = int32(i)+ 9
	n6 := n4+1
	fmt.Println(n4,n5,n6)

	// 第一种方法，使用fmt.Sprintf方法；
	num1 := 99
	var num2 float64 = 23.3254
	var b bool = true
	var mystr byte = 'h'
	var str string 
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%c",mystr)
	fmt.Printf("str type %T str=%q\n",str,str)

	// 第二种反法，strconv 包的方
	var num3 int = 99
	var num4  float64 =23.34634
	var b2 bool = true
	str = strconv.FormatInt(int64(num3),10)
	fmt.Println("===================")
	fmt.Printf("str type %T str=%q\n",str,str) 

	str = strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n",str,str)



}