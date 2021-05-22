package abstract_interface
//message manage
type AMsgHandle interface{
	//do router method
	DoMsgHandle(request ARequest)
	//router specific logic
	AddRouter(msgID uint32,router ARouter)
}