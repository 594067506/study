package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"study/learn"
)


//第二周作业
//业务调用方
func main()  {
	u:=learn.NewUser()
	//查询一行

	err:=learn.QueryOneRow(u)
	switch err {
	case nil:
		fmt.Printf("查询成功：%v",u )
	case sql.ErrNoRows:
		//业务逻辑
		fmt.Println("匹配成功")
	default:
		fmt.Printf("原始错误-调用main函数出错 \n")
		fmt.Printf( "%+v \n", errors.WithStack(err))
	}
	return


}


