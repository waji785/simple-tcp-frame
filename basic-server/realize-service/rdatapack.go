package realize_service

import (
	"bytes"
	"encoding/binary"
	"errors"
	"simple-farme/basic-server/abstract-interface"
	"simple-farme/basic-server/utils"
)

type DataPack struct{

}
func NewDataPack() *DataPack{
	return &DataPack{}
}
//get head length
func(dp *DataPack) GetHeadLen() uint32{
	//Datalen unit32(4byte)+ID uint32(4byte)
	return 8
}
//pack
func(dp *DataPack) Pack(msg abstract_interface.AMessage)([]byte,error){
	//make a buffer to saver byte
	dataBuffer:=bytes.NewBuffer([]byte{})
	//write must be in order
	//write in datalen
	if err :=binary.Write(dataBuffer,binary.LittleEndian,msg.GetMsLen());err !=nil{
		return nil,err
	}
	//write in msgid
	if err :=binary.Write(dataBuffer,binary.LittleEndian,msg.GetMsgId());err !=nil{
		return nil,err
	}
	//write in data
	if err :=binary.Write(dataBuffer,binary.LittleEndian,msg.GetData());err !=nil{
		return nil,err
	}
	return dataBuffer.Bytes(),nil
}
//unpack
//read head first,and then read data
func(dp *DataPack) Unpack(binaryData []byte)(abstract_interface.AMessage,error){
	//make a ioReader
	dataBuff :=bytes.NewReader(binaryData)
	//Unzip head only
	msg :=&Message{}
	//read datalen
	if err :=binary.Read(dataBuff,binary.LittleEndian,&msg.DataLen);err !=nil{
		return nil,err
	}
	//read msgid
	if err :=binary.Read(dataBuff,binary.LittleEndian,&msg.Id);err !=nil{
		return nil,err
	}
	if (utils.GlobalObject.MaxPackageSize>0&&msg.DataLen>utils.GlobalObject.MaxPackageSize) {
		return nil,errors.New("too large msg data")
	}
	return msg,nil
}