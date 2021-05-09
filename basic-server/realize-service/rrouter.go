package realize_service

import "simple-farme/basic-server/abstract-interface"

//base-router,Inheritance base class
type BaseRouter struct{

}
//before conn,Hook func
func (br *BaseRouter) PreHander(request abstract_interface.ARequest){
	//void
}
//method to deal with conn
func (br *BaseRouter) Hander(request abstract_interface.ARequest){
	//void
}
//after conn
func (br *BaseRouter) PostHander(request abstract_interface.ARequest){
	//void
}