package abstract_interface

//interface of request
//packaging request data andconn to interface
type ARequest interface {
	//get current connection
	GetConnection() AConnection
	//get request data
	GetData() []byte
	//get message id
	GetMsgID() uint32
}
