package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	beego_context "github.com/beego/beego/v2/server/web/context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net"
	"study/Grpc/proto"
)

const(
	Address = "127.0.0.1:8801"   // 监听的端口
)

type Greeter struct {
	proto.UnimplementedGoodsServer
}

//实现proto中定义的hello方法
func (greeter *Greeter) Hello(ctx context.Context, request *proto.Request) (*proto.Response, error){
	fmt.Println("get client info, name is: ", request.Name)
	response_msg := "Hello " + request.Name
	resp:=&proto.Response{}
	resp.Msg = response_msg
	resp.Code = 10
	info := map[string]interface{}{
		"name":"guolong",
		"age":10,
	}
	resp.Info,_= json.Marshal(info)
	return resp , nil
}

func RpcServer()  {
	service := Greeter{}
	conn, err := net.Listen("tcp", Address)   // tcp监听端口
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("tcp listening...")
	defer conn.Close()

	server := grpc.NewServer()
	proto.RegisterGoodsServer(server, &service)  // 将tcp服务于grpc进行绑定

	fmt.Println("server serve.....")

	if err := server.Serve(conn); err != nil{
		fmt.Println(err)
		return
	}

}


//gin 的http接口
func HttpServer()  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9000")
}

//beego启动的http服务
func HttpServerForBeego()  {

	type BeeResp struct  {
		Code  string
		Msg  string
	}

	resp:=BeeResp{
		Code: "0",
		Msg: "success",
	}

	//resp_byte,_:=json.Marshal(resp)

	web.Post("/pingb",func(ctx *beego_context.Context){
		//ctx.Output.Body(resp_byte)
		ctx.Output.JSON(resp,true,true)
	})
	web.Run()
}


func main()  {
	RpcServer()
}

