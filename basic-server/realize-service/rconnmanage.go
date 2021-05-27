package realize_service

import (
	"errors"
	"fmt"
	abstract_interface "simple-farme/basic-server/abstract-interface"
	"sync"
)

type ConnManager struct{
	//manage collection
	connections map[uint32]abstract_interface.AConnection
	//read - write lock
	connLock sync.RWMutex
}
func NewConnManager() *ConnManager{
	return &ConnManager{
		connections:make(map[uint32] abstract_interface.AConnection),
	}
}
//add conn
func (ConnManager *ConnManager) Add(conn abstract_interface.AConnection){
	//protect map,add lock
	ConnManager.connLock.Lock()
	defer ConnManager.connLock.Unlock()
	//put conn into ConnManager
	ConnManager.connections[conn.GetConnID()]=conn
	fmt.Println("conID =",conn.GetConnID()," add to ConnManager successfully:conn num=",ConnManager.Len())
}
//remove conn
func (ConnManager *ConnManager) Remove(conn abstract_interface.AConnection){
	//protect map,add lock
	ConnManager.connLock.Lock()
	defer ConnManager.connLock.Unlock()
	fmt.Println("conID =",conn.GetConnID(),"remove from  ConnManager successfully:conn num=",ConnManager.Len())
}
func (ConnManager *ConnManager) Get(connID uint32)(abstract_interface.AConnection,error){
	//protect map,add lock
	ConnManager.connLock.RLock()
	defer ConnManager.connLock.RUnlock()
	if conn,ok:=ConnManager.connections[connID];ok{
		return conn,nil
	}else{
		return nil,errors.New("connection not POUND!")
	}
}
//conn nums
func (ConnManager *ConnManager) Len() int{
	return len(ConnManager.connections)
}
//terminate conn
func (ConnManager *ConnManager) ClearConn(){
	//protect map,add lock
	ConnManager.connLock.Lock()
	defer ConnManager.connLock.Unlock()
	//remove conn and stop
	for connID,conn:=range ConnManager.connections{
		conn.Stop()
		delete(ConnManager.connections,connID)
	}
	fmt.Println("Clear All connection successfully!conn mun=",ConnManager.Len())
}

