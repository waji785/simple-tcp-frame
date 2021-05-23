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
type Ping2Router struct{
	realize_service.BaseRouter
}

//test Handle
func (this *PingRouter) Handle(request abstract_interface.ARequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("recv from client:msgID=",request.GetMsgID(),",data=",string((request.GetData())))
	err :=request.GetConnection().SendMsg(0,[]byte("ping...."))
	if err !=nil{
		println(err)
	}
}
//test Handle
func (this *Ping2Router) Handle(request abstract_interface.ARequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("recv from client:msgID=",request.GetMsgID(),",data=",string((request.GetData())))
	err :=request.GetConnection().SendMsg(201,[]byte("ping2...."))
	if err !=nil{
		println(err)
	}
}

func main() {
	//使用API，创建句柄
	s := realize_service.NewServer("demoV0.6")
	//添加一个自定义router
	s.AddRouter(0,&PingRouter{})
	s.AddRouter(1,&Ping2Router{})
	//启动server
	s.Run()
}
