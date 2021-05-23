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
			fmt.Println("listenerr", err)
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
			//bind with conn to get linked moduel
			dealConn := NewConnection(conn, cid, s.MsgHandle)
			cid++
			go dealConn.Start()
		}

	}()

}
func (s *Server) Stop() {
	//TODO malloc
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

//initialize server
func NewServer(name string) abstract_interface.AServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		MsgHandle:NewMsgHandle(),
	}
	return s
}
