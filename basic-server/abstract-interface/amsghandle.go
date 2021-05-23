package abstract_interface
//message manage
type AMsgHandle interface{
	//do router method
	DoMsgHandle(request ARequest)
	//router specific logic
	AddRouter(msgID uint32,router ARouter)
	//statr workpool
	StartWorkerPool()
	//give msg to mq,by worker process
	SendMsgToTaskQueue(request ARequest)
}