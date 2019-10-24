package main

/* 在Go中，接口是一组方法签名。当类型为接口中的所有方法提供定义时，它被称为实现接口*/
import "fmt"

type Inter interface {
	Get() int
	Set(int)
}

type Stud struct {
	age   int
	name  string
	phone string
}

type Human struct {
	Stud
	school string
}

type Employee struct {
	Human
	company string
	money   float32
}

func (c Stud) Get() int {
	return c.age
}

func (c *Stud) Set(age int) {
	c.age = age
}

func (h Human) SayHi() {
	fmt.Print("hi,I am %s,you can call me on %s\n", h.name, h.phone)
}

// Employee 重写了Sayhi 方法；
func (e Employee) SayHi() {
	fmt.Print("hi I am %s, i work at %s,call me on %s\n", e.name, e.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("la la la la la ....", lyrics)
}

func testInter(i Inter) {
	i.Set(18)
	fmt.Println(i.Get())
}

// interface Men被human student 和 employee 时限，因为这个三个类型都实现了这俩个方法；
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	fmt.Println("interface test...........")
	s := Stud{}
	testInter(&s)

	mike := Human{Stud{18, "Mike", "222-222-xxxx"}, "攀枝花学院"}

	var i Men
	fmt.Println("this is mike a human..")
	i = mike
	i.SayHi()
	i.Sing("we are the world!")

}
