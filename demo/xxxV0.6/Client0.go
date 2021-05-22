package main

import (
	"fmt"
	"io"
	"net"
	realize_service "simple-farme/basic-server/realize-service"
	"time"
)

//模拟客户端
func main() {
	fmt.Println("client0 start")
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8554")
	if err != nil {
		fmt.Println("client start err", err)
		return
	}
	for {
		dp:=realize_service.NewDataPack()
		binaryMsg,err:=dp.Pack(realize_service.NewMsgPackage(0,[]byte("test0 message")))
		if err !=nil{
			fmt.Println("pack error:",err)
			return
		}
		if _,err :=conn.Write(binaryMsg);err!=nil{
			fmt.Println("write error",err)
			return
		}

		binaryHead:=make([]byte,dp.GetHeadLen())
		if _,err:=io.ReadFull(conn,binaryHead);err!=nil{
			fmt.Println("read head error",err)
			break
		}
		msgHead,err :=dp.Unpack(binaryHead)
		if err!=nil{
			fmt.Println("client unpack msghead error",err)
			break
		}
		if msgHead.GetMsgLen()>0{
			msg:=msgHead.(*realize_service.Message)
			msg.Data=make([]byte,msg.GetMsgLen())
			if _,err:=io.ReadFull(conn,msg.Data);err !=nil{
				fmt.Println("read nsg data error",err)
				return
			}
			fmt.Println("recv msg: id=",msg.Id,"len=",msg.DataLen,"data=",msg.Data)
		}
		time.Sleep(1 * time.Second)
	}

}
