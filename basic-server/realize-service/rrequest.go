package realize_service

import "simple-farme/basic-server/abstract-interface"

type Request struct {
	//current connection
	conn abstract_interface.AConnection
	//request data
	data []byte
}

//get current connection
func (r *Request) GetConnection() abstract_interface.AConnection {
	return r.conn
}

//get request data
func (r *Request) GetData() []byte {
	return r.data
}
