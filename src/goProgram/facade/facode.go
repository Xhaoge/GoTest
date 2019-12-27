package main

import "fmt"

func NewApi() API {
	return &allAPI{
		A:newAapi(),
		B:newBapi(),
	}
}

type API interface{
	test() string
	nima()
}

type allAPI struct{
	A aApiMoudle
	B bApiMoudle
}

func (a *allAPI)test() string {
	return "this is api test()"
}

type aApiMoudle interface{
	Say() string
}

type bApiMoudle interface{
	Hello() string
}

type aAPI struct{}

type bAPI struct{}

func (n *aAPI)Say()string{
	return "this is aapi say()"
}

func (m *bAPI)Hello()string{
	return "this is bapi hello()"
}

func newAapi() aApiMoudle{
	return &aAPI{}
}

func newBapi() bApiMoudle{
	return &bAPI{}
}

func (a *allAPI)nima(){
	fmt.Println("this is nima() interface")
	a.A.Say()
	a.B.Hello()
}

func main() {
	fmt.Println("this facade test")
	inter1 := newAapi()
	inter2 := newBapi()
	api1 := allAPI{inter1,inter2}
	api1.nima()
}