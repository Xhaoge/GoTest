package main
/* 方法 test方法只是一个函数，它带有一个特殊的接收器类型，它是在func关键字和方法名之间编写的。
   接收器可以是struct类型或非struct类型。接收方可以在方法内部访问*/
import ("fmt"
		"math")

type Employee struct {
	name string
	salary int
	currency string
}
// displaySalary() method has Employee as the receiver type
func (e Employee) displaySalary(){
	fmt.Printf("salary of %s is %d %s\n",e.name,e.salary,e.currency)
}

// 定义几种图形的面积，长方形和圆；
type Rectangle struct{
	width,height float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) area() float64{
	return r.width * r.height
}

func (r Circle) area() float64 {
	return r.radius * r.radius *math.Pi
}

// method 是可以继承的，如果匿名字段实现了一个method，那么包含这个匿名字段的额struct也能调用该method；
type Human struct{
	name string
	age int
	phone string
}

type Student struct{
	Human //匿名字段
	school string
}

type Employeefrom struct{
	Human
	company string
}

func (h *Human) SayHi(){
	fmt.Printf("Hi,I am %s you can call me on %s\n",h.name,h.phone)
}

/*Employeefrom的method重写Human的method
  - 方法是可以继承和重写的
  - 存在继承关系时，按照就近原则，进行调用
*/
func (e *Employeefrom) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main(){
	fmt.Println("xhaoge.............")
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