package abstract_interface

//define a service interface
type AServer interface {
	//start
	Start()
	//stop
	Stop()
	//run
	Run()
	//router,support of server
	AddRouter(msgID uint32,router ARouter)
	//get the manager of server
	GetConnectionManager() AConnManager
	//register OnConnStart
	SetOnConnStart(func(manager AConnection))
	//register OnConnStop
	SetOnConnStop(func(manager AConnection))
	//call OnConnStart
	CallOnConnStart(manager AConnection)
	//call OnConnStop
	CallOnConnStop(manager AConnection)
}
