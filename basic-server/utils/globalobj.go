package utils

import (
	"encoding/json"
	"io/ioutil"
	"simple-farme/basic-server/abstract-interface"
)

//stores global arguments
//configured by the user with JSON
type GlobalObj struct {
	//global obj
	TcpServer abstract_interface.AServer
	//IP
	Host string
	//port
	TcpPort int
	//server name
	Name string
	//framework version
	Version string
	//Max connection
	MaxConn int
	//MAX data package
	MaxPackageSize uint32
}

//define global obj
var GlobalObject *GlobalObj

//reload from json
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("")
	if err != nil {
		panic(err)
	}
	//resolve to struct
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
func init() {
	//default value
	GlobalObject = &GlobalObj{
		Name:           "xxx",
		Version:        "v0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	//load from json
	GlobalObject.Reload()
}
