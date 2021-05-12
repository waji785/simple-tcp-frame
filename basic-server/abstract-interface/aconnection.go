package abstract_interface

import "net"

//define the interface of linked mod
type AConnection interface {
	//start connection
	Start()
	//stop connection
	Stop()
	//get conn from socket
	GetTCPConnection() *net.TCPConn
	//get connection ID
	GetConnID() uint32
	//get TCP status
	RemoteAddr() net.Addr
	//post data to client
	Send(data []byte) error
}

//todo sth
type HandleFunc func(*net.TCPConn, []byte, int) error
