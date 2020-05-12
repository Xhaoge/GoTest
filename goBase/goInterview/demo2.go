package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type people struct {
	Name string `json:name`
}

func (p *people) String() string {
	return fmt.Sprint("print: %v", p.Name)
}

func testChan() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("CLOSE")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

func testMap() {
	m := map[string]*people{"xxx": {"xaoge"}}
	m["xxx"].Name = "zhangsan"
	fmt.Println("m.studnet name: ", m["xxx"].Name)
}

type query func(string) string

func exec(name string, vs ...query) []string {
	ch := make(chan string, 4)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	//time.Sleep(time.Second * 1)
	res := []string{}
	go func() {
		for k := range ch {
			res = append(res, k)
			fmt.Println("res: ", res)
		}
	}()
	time.Sleep(time.Second * 5)
	close(ch)
	// fmt.Println("ch: ", <-ch)
	fmt.Println("res: ", res)
	return res
}

func testQuery() {
	go func() {
		fmt.Println("youzai zhixingma")
		ret := exec("111", func(n string) string {
			return n + "func1"
		}, func(n string) string {
			return n + "func2"
		}, func(n string) string {
			return n + "func3"
		}, func(n string) string {
			return n + "func4"
		})
		fmt.Println("ret: ", ret)
	}()
}

func main() {
	js := `{
		"Name":"xhoage"
	}`
	var p people
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Println("people: ", p)
	pp := &people{
		Name: "xhaoge",
	}
	pp.String()
	//testChan()
	testQuery()
	// time.Sleep(time.Second * 10)
	testMap()
}
