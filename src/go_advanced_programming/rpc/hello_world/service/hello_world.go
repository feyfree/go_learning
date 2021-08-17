package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//func main() {
//	err := rpc.RegisterName("HelloService", new(HelloService))
//	if err != nil {
//		return
//	}
//
//	listener, err := net.Listen("tcp", ":1234")
//	if err != nil {
//		log.Fatal("ListenTCP error:", err)
//	}
//
//	conn, err := listener.Accept()
//	if err != nil {
//		log.Fatal("Accept error:", err)
//	}
//
//	rpc.ServeConn(conn)
//}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
