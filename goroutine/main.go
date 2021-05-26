package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
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

//这个例子 演示了 在外面设置一个stop信号来控制子携程的退出
func TestGoroutine()  {
	//启动一一个chan在外面控制chan的 正个生命周期
	var stop =make(chan struct{}) //无缓冲的一个管道

	go func() {
		for  {
			select {
			case <-stop:
				fmt.Printf("已收到停止信号,携程退出~~~ /n")  //   收到停止信号退出协程
				return
			default:
				fmt.Printf("默认的逻辑== \n")  		 //   这里可以做真正的逻辑
				time.Sleep(time.Second*1)
			}
		}
	}()

	time.Sleep(time.Second*10) //十秒后发送关闭信号
	fmt.Printf("发送关闭信号= \n")
	stop <- struct{}{}
	//close(stop)

	select {}
}


//测试 errorgroup ，比 waitgroup 多了一个err的传递,当携程返回错误的时候抛出来
func TextErrGroup()  {
	var g errgroup.Group

	var urls = []string{
		"https://www.baidu.com/",
		"https://mp.weixin.qq.com/",
		"https://www.sina.com.cn/",
	}


	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // 循环定义个那个url是一个复用的变量,内存地址是一个,这里如果不重新定义的话，携程fork完后再去找到url地址去拿val都是最后一个

		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("请求成功 url：%s ",url)
				resp.Body.Close()
			}
			return err
		})
	}

	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs:")
	}else {
		fmt.Println("请求报错:"+err.Error())
	}

}