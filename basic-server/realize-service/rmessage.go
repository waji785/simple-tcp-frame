package realize_service

type Message struct {
	Id      uint32 //message id
	DataLen uint32 //message length
	Data    []byte //message region
}
//get message id
func (m *Message) GetMsgId() uint32{
	return m.Id
}
//get message length
func (m *Message) GetMsLen() uint32{
	return  m.DataLen
}
//get data
func (m *Message) GetData() []byte{
	return  m.Data
}
//set data
func (m *Message) SetData(data []byte){
	m.Data=data
}
//set message id
func (m *Message) SetMsgId(id uint32){
	m.Id=id
}
//set data length
func (m *Message) SetDataLen(len uint32){
	m.DataLen=len
}
