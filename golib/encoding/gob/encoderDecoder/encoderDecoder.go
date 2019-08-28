// This example transmits a value that implements the custom
// encoding and decoding methods.
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// The Vector type has unexported fields, which the package cannnot access.
// We therefore write a BinaryMarshal/BinaryUnmarshal method pair to allow us
// to send and receive the type with the gob package.
// These interfaces are defined in the "encoding" packages.
// We could equivalently use the locally defined GobEncode/GobDecoder interfaces.
type Vector struct {
	x, y, z int
}

func (v Vector) MarshalBinary() ([]byte, error) {
	// A simple encoding: plain text.
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

// UnmarshalBinary modified the receiver.
// so it must take a pointer receiver.
func (v *Vector) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}

// Transmit a value that implements the custom encoding and decoding methods.
func main() {
	var network bytes.Buffer // Stand-in for the network.

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("Encode: ", err)
	}

	// Create an decoder and receive a value.
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("Decode: ", err)
	}
	fmt.Println(v)
}
