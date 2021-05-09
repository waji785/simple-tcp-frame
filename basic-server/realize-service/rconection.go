package realize_service

import (
	"fmt"
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
	//router
	Router abstract_interface.ARouter
}
func NewConnection(conn *net.TCPConn,connID uint32,router abstract_interface.ARouter) *Connection{
	c:=&Connection{
		Conn: conn,
		ConnID:connID,
		Router: router,
		isClosed: false,
		ExitChan: make(chan bool,1),
	}
	return c
}
//do sth
func (c *Connection) StartReader(){
	fmt.Println("reader goroutine is running")
	defer fmt.Println("connID=",c.ConnID,"Reader is exit, remote addr is",c.RemoteAddr().String())
	defer c.Stop()
	for {
		//read data buf,maxbuf 512byte
		buf:=make([]byte,512)
		_,err:=c.Conn.Read(buf)
		if err !=nil{
			fmt.Println("recv buf err",err)
			continue
		}
		//get request data
		req :=Request{
			conn:c,
			data:buf,
		}
		//register router
		go func(request abstract_interface.ARequest) {
			//use router
			c.Router.PreHander(request)
			c.Router.Hander(request)
			c.Router.PostHander(request)
		}(&req)
	}
}
//start connection
func (c *Connection) Start(){
	fmt.Println("conn start... connID=",c.ConnID)
	//todo sth
	go c.StartReader()
}
//stop connection
func (c *Connection) Stop(){
	fmt.Println("conn stop... connID=",c.ConnID)
	if c.isClosed==true{
		return
	}
	c.isClosed=true
	//close socket
	c.Conn.Close()
	close(c.ExitChan)
}
//get conn from socket
func (c *Connection) GetTCPConnnection() *net.TCPConn{
	return c.Conn
}
//get connection ID
func (c *Connection) GetConnID() uint32{
	return c.ConnID
}
//get TCP status
func (c *Connection) RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()

}
//post data to client
func (c *Connection) Send(data []byte) error{
	return nil
}
