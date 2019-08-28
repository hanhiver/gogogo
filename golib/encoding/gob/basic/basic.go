// This example shows the basic usage of the package:
// Create an encoder,
// Transmit some values,
// Receive them with a decoder.
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	// Initialize the encoder and decoder.
	// Normally, enc and dec would be bound to network connections,
	// and encoder and decoder would run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection.
	enc := gob.NewEncoder(&network) // will write to network.
	dec := gob.NewDecoder(&network) // will read from network.

	// Encode (send) some values.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("Encoder error: ", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("Encoder error: ", err)
	}

	// Decode (receive) and print the values.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("Dncoder error 1: ", err)
	}
	fmt.Printf("%q: {%d, %d} \n", q.Name, *q.X, *q.Y)

	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("Dncoder error 2: ", err)
	}
	fmt.Printf("%q: {%d, %d} \n", q.Name, *q.X, *q.Y)
}
