package main

import (
	"context"
	"fmt"
	"time"
)



func main()  {
	TestCancelContext()
	//TestValueContext()
}

// 测试context 的 可取消的 context 和 带过期时间的 context
//可取消context:一个函数中可能开启多个goroutine来处理 不同的业务，但是我们要可以控制整个生命周期，主方法中可以调用cancle() 子goroutine中收到取消信号后退出
//过期时间的 context:同可取消context 类似，但是多了一个过期时间，方时间过期后， 子goroutine中也会收到取消信号并退出

func TestCancelContext() {

	ctx := context.Background()

	//ctx1,cancle :=context.WithCancel(ctx,time.Second*10)
	ctx1,cancle :=context.WithTimeout(ctx,time.Second*11)

	go DoSth(ctx1)

	time.Sleep(time.Second*10)
	fmt.Println("发送停止信号=")
	cancle()

	select {}
}

func DoSth(ctx context.Context) {
	i:=1
	for  {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx 取消 携程退出 ")
			return
		default:
			fmt.Println("执行业务逻辑：",i)
			i++
			time.Sleep(time.Second*1)
		}
	}

}



//测试VauleContext：这个的内部 就是增加了一个key 和 val，如果是嵌套的话 会有一个链表来管理所有的子代关系，如例子中
//ctx2 继承了 ctx1 在DoSth1中查找 userid的时候 ctx2中没有userid 会沿着关系链一直向上查询，直到查询到第一个含有userid的val为止

func TestValueContext() {
	ctx := context.Background()
	ctx1 :=context.WithValue(ctx,"userid",1000)
	ctx2,cancle :=context.WithTimeout(ctx1,time.Second*10)
	defer cancle()

	go DoSth1(ctx2)


	select {}

}

func DoSth1(ctx context.Context) {

	for  {
		select {
		case <-ctx.Done():
			fmt.Println("ctx超时，goroutine退出==")
			return
		default:
			fmt.Println(ctx.Value("userid"))
		}

	}

}