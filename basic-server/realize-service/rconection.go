package realize_service

import (
	"errors"
	"fmt"
	"io"
	"net"
	"simple-farme/basic-server/abstract-interface"
)

type Connection struct {
	//socket TCP
	Conn *net.TCPConn
	//ID
	ConnID uint32
	//status
	isClosed bool
	//bind api
	handleAPI abstract_interface.HandleFunc
	//stop channel
	ExitChan chan bool
	//buff ,used for message communication between go
	msgChan chan []byte
	//router
	MsgHandle abstract_interface.AMsgHandle
}

func NewConnection(conn *net.TCPConn, connID uint32, msgHandler abstract_interface.AMsgHandle) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		MsgHandle:   msgHandler,
		isClosed: false,
		msgChan: make(chan []byte),
		ExitChan: make(chan bool, 1),
	}
	return c
}

//do sth
func (c *Connection) StartReader() {
	fmt.Println("reader goroutine is running")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()
	for {
		dp:=NewDataPack()
		//read message head 8byte
		headData:=make([]byte,dp.GetHeadLen())
		if _,err :=io.ReadFull(c.GetTCPConnection(),headData);err!=nil{
			fmt.Println("read msg head err",err)
			break
		}
		//unpack get message id and datalen
		msg,err:=dp.Unpack(headData)
		if err!=nil{
			fmt.Println("unpack err",err)
			break
		}
		//put dtat in message.Data
		var data []byte
		if msg.GetMsgLen()>0{
			data =make([]byte,msg.GetMsgLen())
			if _,err :=io.ReadFull(c.GetTCPConnection(),data);err!=nil{
				fmt.Println("read msg data err",err)
				break
			}
		}
		msg.SetData(data)
		//get request data
		req := Request{
			conn: c,
			msg: msg,
		}
		//register router
		go c.MsgHandle.DoMsgHandle(&req)
	}
}
func (c *Connection) StartWriter(){
	fmt.Println("Write Gortine is running")
	defer fmt.Println(c.RemoteAddr().String(),"[conn Write exit]")
	for{
		select {
		case data:=<-c.msgChan:
			if _,err:=c.Conn.Write(data);err !=nil{
				fmt.Println("Send data error",err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}
//send msg method,pack data and send
func (c *Connection) SendMsg(msgId uint32, data []byte) error{
	if c.isClosed==true{
		return errors.New("Connection closed when send msg")
	}
	//pack data
	dp:=NewDataPack()
	binaryMsg,err:=dp.Pack(NewMsgPackage(msgId,data))
	if err !=nil{
		fmt.Println("Pack error msg id=",msgId)
		return errors.New("Pack error msg")
	}
	c.msgChan<-binaryMsg
	return nil

}

//start connection
func (c *Connection) Start() {
	fmt.Println("conn start... connID=", c.ConnID)
	//todo sth
	go c.StartReader()
	go c.StartWriter()
}

//stop connection
func (c *Connection) Stop() {
	fmt.Println("conn stop... connID=", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	//close socket
	c.Conn.Close()
	close(c.ExitChan)
}

//get conn from socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//get connection ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//get TCP status
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()

}

