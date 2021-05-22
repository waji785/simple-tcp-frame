package realize_service

import (
	"fmt"
	"simple-farme/basic-server/abstract-interface"
	"strconv"
)

type MsgHandle struct{
	//saving that handing method for every id
	Apis map[uint32] abstract_interface.ARouter
}
func NewMsgHandle() *MsgHandle{
	return &MsgHandle{
		Apis: make(map[uint32]abstract_interface.ARouter),
	}
}
//do router method
func (mh *MsgHandle)DoMsgHandle(request abstract_interface.ARequest){
	//select msg id
	handler,ok:=mh.Apis[request.GetMsgID()]
	if !ok{
		fmt.Println("api msgID",request.GetMsgID(),"is NOT POUND! Need Register!")
	}
	//select service
	handler.PreHander(request)
	handler.Hander(request)
	handler.PostHander(request)
}
//router specific logic
func (mh *MsgHandle)AddRouter(msgID uint32,router abstract_interface.ARouter){
	//judge if method exist
	if _,ok:=mh.Apis[msgID];ok{
		panic("repeat api,msgID="+strconv.Itoa(int(msgID)))
	}
	//if not exist then add
	mh.Apis[msgID]=router
	fmt.Println("Add api MsgID=",msgID,"success!")
}