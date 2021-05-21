package abstract_interface
//the interface that encapsulates the message
type AMessage interface {
	//get message id
	GetMsgId() uint32
	//get message length
	GetMsgLen() uint32
	//get data
	GetData() []byte
	//set data
	SetData([]byte)
	//set message id
	SetMsgId(uint32)
	//set data length
	SetDataLen(uint32)
}