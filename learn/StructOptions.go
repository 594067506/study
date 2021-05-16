package learn

import "fmt"

type Server struct {
	network string
	address string
}

type ServerOption func(*Server)

func Network(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

func Addr(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

func SetServer(opts ...ServerOption)  {

	//先做必要的初始化
	s:=&Server{
		network: "192.168.1.1",
		address: "静安",
	}

	fmt.Println("fi=====")
	fmt.Println(s.address)
	fmt.Println(s.network)

	for   _,o:= range opts {
		o(s)
	}

	fmt.Println("tw=====")
	fmt.Println(s.address)
	fmt.Println(s.network)
}

//options思想调用
func Do()  {
	SetServer(Network("111"),Addr("松江区"))
}
