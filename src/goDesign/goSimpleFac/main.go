 package main

 import (
 	"fmt"
 )

// api is interface
type API interface {
	Say(name string) string
}

// newapi return api interface bu type
func NewApi(t int) API {
	if t == 1 {
		return &hiAPI{}
	}else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

//hiAPI is one of API implement
type hiAPI struct{
	Name string
}

// Say hi to name

func (h *hiAPI) Say(name string) string {
	h.Name = name
	return fmt.Sprintf("hi ,%v,hiAPI say()",h.Name)
}

// helloAPI is another API implement
type helloAPI struct{
	Work string
}

// Say hello to name
func (e *helloAPI) Say(name string) string {
	e.Work = name
	return fmt.Sprintf("hi ,%v,helloAPI say()",e.Work)
}


 func main (){
 	fmt.Println("this is go simple case!")
 	api1 := NewApi(1)
 	api1.Say("xhaoge")
 	api2 := NewApi(2)
 	api2.Say("name")
 	fmt.Println(api1)
 	fmt.Println(api2)
 }