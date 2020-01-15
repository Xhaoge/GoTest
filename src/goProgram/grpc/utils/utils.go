package utils


import (
	"fmt"
	"Pro_golang/Golang/config"
	"reflect"
	//"strconv"
	//"Pro_golang/Golang/TestCase/Search"
	"Pro_golang/Golang/TestCase/SearchCase"
)

func Utils() {
	fmt.Println("this is utils package")
}

func CaseRunMethod interface{
	RunCaseProcess()
}

func init(){
	fmt.Println("this is utils package")
	// var regStruct map[string]interface{}
	// regStruct = make(map[string]interface{})
	// fmt.Println("regStruct:",regStruct)
	
	for i,c := range config.CaseConfig.Cases {
		fmt.Println(i,c.Casename)
		cc := SearchCase.SearchCase_0002{}
		if c.Casename == reflect.TypeOf(&cc).Elem().Name(){
			// cc := SearchCase.SearchCase_0002
			cc.RunCaseProcess()
		}
	} 
}