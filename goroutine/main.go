package main

import (
	"context"
	"fmt"
	"time"
)


//结构体嵌套inteface 例子  start
type Work interface {
	Do(string) string
}

type  Dog int

func ( p * Dog) Do(s string)  string {
	fmt.Println("dog do ... ",s)
	return s
}

type People struct {
	Name string 
	Age int
}



func ( p * People)Do(s string)  string {
	fmt.Println("people do ... ",s)
	return s
}

type Student struct {
	Work  //嵌套interface在，主要是为了解耦合，任意实现了此机构的类型，都可以注入进来，Student就继承了此类型的属性和方法
	StudentNum string
}


func ( p * People)Study()  {
	fmt.Println("Study.....")
}

//var d Dog
//s:= &Student{Work:&d}
//s.Do("ddd")
//结构体嵌套inteface 例子  end



func main() {
	tr := NewTracker()
	// 是否启动goroutine 应该交给调用者
	go tr.Run()
	_ = tr.Event("test1")
	_ = tr.Event( "test2")
	_ = tr.Event("test3")
	time.Sleep(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
	_ = tr.Event(context.Background(), "test4") // close channel后再发 会painc



}

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{ch: make(chan string, 10)}
}

func (t *Tracker) Event(  data string) error {
	select {
	case t.ch <- data:
		return nil
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}