package realize_service

import (
	"fmt"
	"net"
	"simple-farme/basic-server/abstract-interface"
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

}
//start server
func (s *Server) Start()  {
	fmt.Printf("Server Listenner at IP:%s,Port%d,is starting\n",s.IP,s.Port)
	//if it's not use go func,then start() will be always block
	go func() {
		//get a TCP addr
		addr,err:=net.ResolveTCPAddr(s.IPVersion,fmt.Sprint("%s:%d",s.IP,s.Port))
		if err !=nil{
			fmt.Println("addr error",err)
			return
		}
		//listen TCP addr
		listenner,err :=net.ListenTCP(s.IPVersion,addr)
		if err !=nil{
			fmt.Println("listenerr",err)
			return
		}
		fmt.Println("start successfully")
		//block and wait,for read and write
		for{
			conn,err :=listenner.AcceptTCP()
			if err !=nil{
				fmt.Println("accept err",err)
				continue
			}
			//service code
			//callback
			go func() {
				for  {
					buf:=make([]byte,512)
					cnt,err :=conn.Read(buf)
					if err !=nil {
						fmt.Println("recevie buf err",err)
						continue
					}
					//callback
					if _,err :=conn.Write(buf[:cnt]);err!=nil{
						fmt.Println("callback buf err",err)
						continue
					}
				}
			}()
		}

	}()

}
func (s *Server) Stop()  {
	//TODO malloc
}
func (s *Server) Run()  {
	s.Start()
	//TODO do sth

	//block,for do sth
}
//initialize server
func NewServer(name string) abstract_interface.Aserver{
	s:=&Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8554,
	}
	return s
}