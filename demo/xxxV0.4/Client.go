package main

import (
	"fmt"
	"net"
	"time"
)

//模拟客户端
func main()  {
	fmt.Println("client start")
	time.Sleep(1*time.Second)
	conn,err :=net.Dial("tcp","127.0.0.1:8554")
	if err !=nil{
		fmt.Println("client start err",err)
		return
	}
	for{
		_,err :=conn.Write([]byte("hello"))
		if err !=nil{
			fmt.Println("write conn err",err)
			return
		}
		buf :=make([]byte,512)
		cnt ,err :=conn.Read(buf)
			fmt.Println("read buf err")
			return

		fmt.Printf("server call back: %s,cnt=%d\n",buf,cnt)
		time.Sleep(1*time.Second)
	}

}
