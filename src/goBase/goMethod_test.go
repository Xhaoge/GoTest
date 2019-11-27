package main

import (
	"testing"
	"fmt"
)

func TestMethod(t *testing.T){
	fmt.Println("testing !")
	//e = Employee{name:"xhaoge",salary:7500,currency:"USD"}
	fmt.Println("xhaoge......这个是test.......")
	emp1 := Employee{
		name:"Sam Adolf",
		salary:7500,
		currency:"CNY",
	}
	emp1.displaySalary()  // Calling displaySalary() method Employee type
	r1 := Rectangle{12,4.5}
	c1 := Circle{9}
	fmt.Println("Area of r1 is : ",r1.area())
	fmt.Println("Area of c1 is : ",c1.area())
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employeefrom{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}