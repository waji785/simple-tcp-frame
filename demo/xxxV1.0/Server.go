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
func DoConnectionBegin(conn abstract_interface.AConnection)  {
	if err:=conn.SendMsg(202,[]byte("DoConnection Begin"));err!=nil{
		fmt.Println(err)
	}
	fmt.Println("set conn")
	conn.SetProperty("name","1")
	conn.SetProperty("id","2")
}
func DoConnectionLost(conn abstract_interface.AConnection)  {
	fmt.Println("conn ID=",conn.GetConnID(),"has lost...")
	if name,err:=conn.GetProperty("name");err==nil{
		fmt.Println("name=",name)
	}
	if id,err:=conn.GetProperty("id");err==nil{
		fmt.Println("id=",id)
	}
	if property,err:=conn.GetProperty("property");err==nil{
		fmt.Println("property=",property)
	}
}

func main() {
	//使用API，创建句柄
	s := realize_service.NewServer("demoV1.0")
	//hook
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	//添加一个自定义router
	s.AddRouter(0,&PingRouter{})
	s.AddRouter(1,&Ping2Router{})
	//启动server
	s.Run()
}
