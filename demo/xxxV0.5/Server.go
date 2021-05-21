package main

import (
	"fmt"
	"simple-farme/basic-server/abstract-interface"
	"simple-farme/basic-server/realize-service"
)

//基于框架开发的服务端
//ping test
type PingRouter struct {
	realize_service.BaseRouter
}

//test PreHandle
func (this *PingRouter) PreHandle(request abstract_interface.ARequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping.."))
	if err != nil {
		fmt.Println("call back ping before error")
	}
}

//test Handle
func (this *PingRouter) Handle(request abstract_interface.ARequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping.."))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

//test PostHandle
func (this *PingRouter) PostHandle(request abstract_interface.ARequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping.."))
	if err != nil {
		fmt.Println("call back ping behind error")
	}
}
func main() {
	//使用API，创建句柄
	s := realize_service.NewServer("demoV0.3")
	//添加一个自定义router
	s.AddRouter(&PingRouter{})
	//启动server
	s.Run()
}
