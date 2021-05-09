package main

import (
	"time"
)

//第二周作业
//业务调用方
//func work2()  {
//	u:=learn.NewUser()
//	//查询一行
//	err:=learn.QueryOneRow(u)
//	switch err {
//	case nil:
//		fmt.Printf("查询成功：%v",u )
//	case sql.ErrNoRows:
//		//业务逻辑
//		fmt.Println("匹配成功")
//	default:
//		fmt.Printf("原始错误-调用main函数出错 \n")
//		fmt.Printf( "%+v \n", errors.WithStack(err))
//	}
//	return
//}


func writeRoutine(test_chan chan int, value int) {
	test_chan <- value

}

func readRoutine(test_chan chan int) {
	time.Sleep(time.Second*5)
	<-test_chan
	return
}
func main()  {
	c := make(chan int)
	go writeRoutine(c, 100)
	readRoutine(c)
	time.Sleep(time.Second*20)
	//readRoutine(c)

}


