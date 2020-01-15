package main

import (
	"fmt"
	"reflect"
)

type person struct{
	Name    string
}

func (p person)test(){
	fmt.Println("test struct")
}

func (p person)Test(n1 ,n2 int)int{
	fmt.Println("Test struct")
	return n1+n2
}

func (p person)Run(){
	fmt.Println("Test struct Run() method")
}



func main(){
	fmt.Println("test")
	m1 := person{"xhaoge"}
	m2 := person{"sunhao"}
	var searchList []interface{}
	searchList =  append(searchList,m1,m2)
	Caselist :=make(map[string]interface{})
	fmt.Println(searchList)
	for _,g := range searchList{
		getType := reflect.TypeOf(g)
		getValue := reflect.ValueOf(g)
		fmt.Println("getType:",reflect.TypeOf(getType))
		knd := getType.Kind()
		fmt.Println("knd:",knd)
		for i:=0;i<getType.NumField();i++{
			field := getType.Field(i)
			value := getValue.Field(i).Interface()
			fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)

		}
		//获取方法 1.先获取interface的reflect type 然后通NumMethod进行遍历
		for i:=0;i<getType.NumMethod(); i++{
			m := getType.Method(i)
			fmt.Println(m.Name, m.Type)
		}

		// var methodlist  []reflect.Value
		// var ff reflect.Value
		// ff = getValue.MethodByName("Test")
		// if !ff.IsValid() {
		// // 如果结构体不存在此方法，输出Panic
		// 	fmt.Println("结构体不存在此方法，输出Panic")
		// }
		// methodlist = append(methodlist,reflect.ValueOf(10))
		// methodlist = append(methodlist,reflect.ValueOf(34))
		// ff.Call(methodlist)
		var methodlist  []reflect.Value
		var ff reflect.Value
		ff = getValue.MethodByName("Run")
		if !ff.IsValid() {
		// 如果结构体不存在此方法，输出Panic
			fmt.Println("结构体不存在此方法，输出Panic")
		}
		ff.Call(methodlist)
		
	}
	
	Caselist["search"] = searchList

	fmt.Println("Caselist",Caselist)
	t := reflect.TypeOf(Caselist["search"]).Kind()
	s := reflect.ValueOf(Caselist["search"])
	fmt.Println(t)
	fmt.Println(s,reflect.TypeOf(s))
	// for _,n := range s{
	// 	c := reflect.TypeOf(n)
	// 	fmt.Println(c)
	// }

}