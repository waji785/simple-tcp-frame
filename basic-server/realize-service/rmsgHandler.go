package realize_service

import (
	"fmt"
	"simple-farme/basic-server/abstract-interface"
	"simple-farme/basic-server/utils"
	"strconv"
)

type MsgHandle struct{
	//saving that handing method for every id
	Apis map[uint32] abstract_interface.ARouter
	//mq
	TaskQueue []chan abstract_interface.ARequest
	//worker pool
	WorkerPoolSize uint32
}
func NewMsgHandle() *MsgHandle{
	return &MsgHandle{
		Apis: make(map[uint32]abstract_interface.ARouter),
		WorkerPoolSize:utils.GlobalObject.WorkerPoolSize,
		TaskQueue: make([]chan abstract_interface.ARequest,utils.GlobalObject.WorkerPoolSize),
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
//start workerpool,only do one time
func (mh *MsgHandle ) StartWorkerPool()  {
	for i:=0;i<int(mh.WorkerPoolSize);i++{
		mh.TaskQueue[i]=make(chan abstract_interface.ARequest,utils.GlobalObject.MaxWorkerTaskLen)
		go mh.StartOneWorker(i,mh.TaskQueue[i])
	}
}
//start worker workflow
func (mh *MsgHandle ) StartOneWorker(workerID int,taskQueue chan abstract_interface.ARequest)  {
	fmt.Println("worker id=",workerID,"is started...")
	for  {
		select{
			case request:=<-taskQueue:
			mh.DoMsgHandle(request)
		}
	}
}
//give msg to mq,by worker process
func (mh *MsgHandle) SendMsgToTaskQueue(request abstract_interface.ARequest){
	//equal distribution
	workerID:=request.GetConnection().GetConnID()%mh.WorkerPoolSize
	fmt.Println("Add ConnID=",request.GetConnection().GetConnID(),
					"request MsgID",request.GetMsgID(),
					"to workerID=",workerID)
	//send msg to the taskQueue
	mh.TaskQueue[workerID]<-request
}