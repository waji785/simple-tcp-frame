package realize_service

import (
	"fmt"
	"net"
	"simple-farme/basic-server/abstract-interface"
	"simple-farme/basic-server/utils"
)

//the realize of server interface,define a server struct
type Server struct {
	//name
	Name string
	//IP.ver
	IPVersion string
	//IP.port
	Port int
	//IP
	IP string
	//router
	MsgHandle abstract_interface.AMsgHandle
	//server conn manager
	ConnManager abstract_interface.AConnManager
	//hook func
	OnConnStart func(conn abstract_interface.AConnection)
	//stop hook before destroy conn
	OnConnStop func(conn abstract_interface.AConnection)

}

//start server
func (s *Server) Start() {
	fmt.Printf("Server Name: %s, listenner at IP: %s,Port:%d is starting",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Sprintf("framework version :%s,Maxconn:%d,MaxPackageSize:%d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	//if it's not use go func,then start() will be always block
	go func() {
		//start mq and workPool
		s.MsgHandle.StartWorkerPool()
		//get a TCP addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprint("%s,%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("addr error", err)
			return
		}
		//listen TCP addr
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listener", err)
			return
		}
		fmt.Println("start successfully")
		var cid uint32
		cid = 0
		//block and wait,for read and write
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("accept err", err)
				continue
			}
			//judge if connections is exceeded
			if s.ConnManager.Len()>=utils.GlobalObject.MaxConn{
				//TODO return error to client
				fmt.Println("Too Many Connection",utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}
			//bind with conn to get linked module
			dealConn := NewConnection(s,conn, cid, s.MsgHandle)
			cid++
			go dealConn.Start()
		}

	}()

}
func (s *Server) Stop() {
	//TODO malloc
	fmt.Println("[Stop]xxx server name",s.Name)
	s.ConnManager.ClearConn()
}
func (s *Server) Run() {
	s.Start()
	//TODO do sth

	//block,for do sth
}
func (s *Server) AddRouter(msgID uint32,router abstract_interface.ARouter) {
	s.MsgHandle.AddRouter(msgID,router)
	fmt.Println("Add Router successfully")
}
func (s *Server) GetConnectionManager() abstract_interface.AConnManager{
	return s.ConnManager
}
//initialize server
func NewServer(name string) abstract_interface.AServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		MsgHandle:NewMsgHandle(),
		ConnManager: NewConnManager(),
	}
	//add conn to connManager
	return s
}
//register OnConnStart
func(s *Server) SetOnConnStart(hookFunc func(connection abstract_interface.AConnection)){
	s.OnConnStart=hookFunc
}
//register OnConnStop
func(s *Server) SetOnConnStop(hookFunc func(connection abstract_interface.AConnection)){
	s.OnConnStop=hookFunc
}
//call OnConnStart
func(s *Server) CallOnConnStart(conn abstract_interface.AConnection){
	if s.OnConnStart!=nil{
		fmt.Println("Call OnConnStart..")
		s.OnConnStart(conn)
	}
}
//call OnConnStop
func(s *Server) CallOnConnStop(conn abstract_interface.AConnection){
	if s.OnConnStop !=nil{
		fmt.Println("Call OnConnStop...")
		s.OnConnStop(conn)
	}
}
