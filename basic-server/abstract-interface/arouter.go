package abstract_interface
//the interface of router
type ARouter interface {
	//before conn,Hook func
	PreHander(request ARequest)
	//method to deal with conn
	Hander(request ARequest)
	//after conn
	PostHander(request ARequest)
}
