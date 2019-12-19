package main

import (
	"fmt"
	"reflect"
)
type SearchCase struct{

}


func main() {
	var SearchCase SearchCase
	fmt.Println(SearchCase)
	cc := reflect.ValueOf(&SearchCase)
	ss := reflect.TypeOf(SearchCase).String()
	fmt.Println(cc)
	fmt.Println(ss) // 返回*main.SearchCase

	dd := reflect.TypeOf(&SearchCase).Elem().Name()
	fmt.Println(dd) // 返回SearchCase 
	aa := reflect.TypeOf(dd)
	fmt.Println(aa) // string
	
}

