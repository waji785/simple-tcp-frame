package abstract_interface
type AConnManager interface {
	//add conn
	Add(conn AConnection)
	//remove conn
	Remove(conn AConnection)
	//get conn
	Get(connID uint32)(AConnection,error)
	//conn nums
	Len() int
	//terminate conn
	ClearConn()
}