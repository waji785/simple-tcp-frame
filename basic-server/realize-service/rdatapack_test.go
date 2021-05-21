package realize_service
//testing class
import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T){
	listenner,err:=net.Listen("tcp","127.0.1:7237")
	if err !=nil{
		fmt.Println("server listen err:",err)
		return
	}
	go func(){
		for{
			conn,err:=listenner.Accept()
			if err !=nil{
				fmt.Println("server accept err:",err)
			}
			go func(conn net.Conn){
				dp:=NewDataPack()
				for{
					//read head
					headData:=make([]byte,dp.GetHeadLen())
					if err !=nil{
						fmt.Println("read head err:",err)
						return
					}
					msgHead,err:=dp.Unpack(headData)
					if err !=nil{
						fmt.Println("server unpack err:",err)
						return
					}
					if msgHead.GetMsgLen()>0{
						//read data
						msg:=msgHead.(*Message)
						msg.Data=make([]byte,msg.GetMsgLen())
						_,err:=io.ReadFull(conn,msg.Data)
						if err !=nil{
							fmt.Println("server unpack data err:",err)
							return
						}
						//finished
						fmt.Println("receive",msg.Id,"data length",msg.DataLen,"data",string(msg.Data))
					}
				}
			}(conn)
		}
	}()
	conn,err:=net.Dial("tco","127.0.0.1:7237")
	if err !=nil{
		fmt.Println("client dial err:",err)
		return
	}
	dp:=NewDataPack()
	msgl:=&Message{
		Id:1,
		DataLen:5,
		Data:[]byte{'1','2','3','4','5'},
	}
	sendDatal,err:=dp.Pack(msgl)
	if err !=nil{
		fmt.Println("client pack msgl err",err)
		return
	}
	msgl2:=&Message{
		Id:2,
		DataLen:8,
		Data:[]byte{'1','2','3','4','5','6','7','8'},
	}
	sendDatal2,err:=dp.Pack(msgl2)
	if err !=nil{
		fmt.Println("client pack msgl err",err)
		return
	}
	sendDatal=append(sendDatal,sendDatal2...)
	conn.Write((sendDatal))
	select{}
}