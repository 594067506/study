package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"study/Grpc/proto"
)

const(
	Address = "127.0.0.1:8801"    // 服务监听的端口
)

func RpcRequest()  {

	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil{
		fmt.Println(err)
	}

	defer conn.Close()
	client := proto.NewGoodsClient(conn)    //  自动生成方法, 因为proto文件中service的名字是hello
	name := "jhonsmith"
	result, err := client.Hello(context.Background(), &proto.Request{Name:name})   // 调用grpc方法, 对服务端进行通讯

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(result)

	fmt.Println("I see, u see. I say ", name, "u say ", result.Msg)
}


func main()  {


}
