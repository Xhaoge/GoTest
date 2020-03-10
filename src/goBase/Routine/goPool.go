package main

import (
	"fmt"
	"time"
)

// 每个任务Task类型，可以抽象成一个函数；
type Task struct {
	f func() error // 一个无参的函数类型
}

//通过一个NewTask 来创建一个task
func NewTask(fx func() error) *Task {
	t := Task{
		f: fx,
	}
	return &t
}

// Task 执行任务的方法
func (t *Task) Execute() {
	t.f() // 调用任务所绑定的函数
}

// 有关协程池的定义及操作
type Pool struct {
	InternalInter chan *Task
	ExternalInter chan *Task
	WorkNum       int
}

// 创建一个协程池；
func NewPool(cap int) *Pool {
	p := Pool{
		InternalInter: make(chan *Task),
		ExternalInter: make(chan *Task),
		WorkNum:       cap,
	}

	return &p
}

// 协程池创建一个worker 并且开始工作
func (p *Pool) worker(workId int) {
	// worker 不断的从internal内部任务中拿任务
	for task := range p.InternalInter {
		//如果拿到任务，则执行task 任务；
		task.Execute()
		fmt.Println("workid:", workId, " 执行任务完毕")
	}
}

// 让协程池pool 开始工作
func (p *Pool) Run() {
	//1.首先根据协程池的worker 数量限定，开启固定的worker.
	for i := 0; i < p.WorkNum; i++ {
		go p.worker(i)
	}
	//2.从external 协程池入口获取外界传递过来的任务；
	for task := range p.ExternalInter {
		p.InternalInter <- task
	}
	//3.执行完毕需要关闭ExternalInter
	//4.执行完毕需要关闭InternalInter
}

func printMyName() error {
	fmt.Println("my name is xxxx")
	return nil
}

func main() {
	fmt.Println("this is pool test")
	// 创建一个task
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	printTask := NewTask(printMyName)

	// 创建一个协程池，最大开启四个worker
	p := NewPool(4)
	go p.ExternalInter <- printTask

	//开一个协程，不断的向pool 输送打印一条时间的task任务；
	go func() {
		for {
			p.ExternalInter <- t
		}
	}()

	// 启动协程池
	p.Run()

}
