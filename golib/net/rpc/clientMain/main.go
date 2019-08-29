package main

import (
	"fmt"
	"gogogo/golib/net/rpc/server"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	// Synchronous call
	args := &server.Args{7, 8}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Arith error:", err)
	}

	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)
}
