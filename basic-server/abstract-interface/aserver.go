package abstract_interface
//define a service interface
type  AServer interface{
	//start
	Start()
	//stop
	Stop()
	//run
	Run()
	//router,support of server
	AddRouter(router ARouter)
}