package main

import (
	"net/rpc"
	"rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

func main() {

	rpc.Register(rpcdemo.Demoservice{})
	listener,err:=net.Listen("tcp",":1234")
	if err!=nil{
		panic(err)
	}
	for  {
		conn,err:=listener.Accept()
		if err!=nil{
			log.Printf("accepterrr:%v",err)
		}
		go jsonrpc.ServeConn(conn)

	}
}
