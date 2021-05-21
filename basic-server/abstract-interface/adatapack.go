package abstract_interface

//the mod of packet and unpacking
//used to deal with the sticky packet problem
//solution is add tcp head
type ADataPack interface{
	//get head length
	GetHeadLen() uint32
	//pack
	Pack(msg AMessage)([]byte,error)
	//unpack
	Unpack([]byte)(AMessage,error)
}