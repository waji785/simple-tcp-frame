package main
import "simple-farme/basic-server/realize-service"
//基于框架开发的服务端
func main(){
	//使用API，创建句柄
	s:=realize_service.NewServer("demoV0.2")
	//启动server
	s.Run()
}