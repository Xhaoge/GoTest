package main

import (
	"fmt"
)

type Task struct {
	f func() error
}

func NewTask(fx func() error) *Task {
	t := Task{
		f: fx,
	}
	return t
}

func (t *Task) Execute() {
	t.f()
}

type Pool struct {
	InternalInter chan Task
	ExternalInter chan Task
	WorkNum       int
}

func (p *Pool) Run() {
	fmt.Println("pool run")
}

func main() {
	fmt.Println("this is pool test")

}
