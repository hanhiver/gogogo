package main

import (
	"gogogo/golib/net/rpc/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	http.Serve(l, nil)
}
