package main

import (
	"fmt"
	"reflect"
)


type Student struct {
	Name	string
	Age 	int
}

func main(){
	var num float64 = 1.23
	fmt.Println("num :",num)
	// 需要操作指针， 获取value 对象,注意必须是指针；
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()
	fmt.Println("类型：",newValue.Type())
	fmt.Println("是否可以重新赋值：",newValue.CanSet())

	// 重新赋值；
	newValue.SetFloat(3.24)
	fmt.Println(num)
	// 如果valueof 的参数不是指针；
	value := reflect.ValueOf(num)
	fmt.Println(value.CanSet())

	value.Elem() // 如果非指针 直接报panic；
	s1 := Student{"wangsan",34}
	p1 := &s1


	// 改变数值；
	va :=reflect.ValueOf(p1)
	if va.Kind() == reflect.Ptr{
		val := va.Elem()
		fmt.Println(val.CanSet())
		
	}


}