package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080",
		http.FileServer(http.Dir("/Users/dhan/go/src/gogogo/golib/net/http/fileServer"))))
}
