package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		// wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// echo all incoming data.
			io.Copy(c, c)
			// shut down the connection.
			c.Close()
		}(conn)
	}
}
