package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Hypotenuse() float64 {
	return math.Hypot(float64(p.X), float64(p.Y))
}

type Pythagoras interface {
	Hypotenuse() float64
}

// interfaceEncoder encodes the interface value into the encoder.
func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	err := enc.Encode(&p)
	if err != nil {
		log.Fatal("Encode: ", err)
	}
}

// interfaceEncoder dncodes the interface value into the encoder.
func interfaceDecode(dec *gob.Decoder) Pythagoras {
	var p Pythagoras
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("Decode: ", err)
	}
	return p
}

// This example shows how to encode an interface value.
// The key distinction from regular types is to register the concrete type that
// implements the interface.
func main() {
	var network bytes.Buffer

	// We must register the concrete type for the encoder and decoder.
	// On each end, this tells the engine which concrete type is being sent
	// that implements the interface.
	gob.Register(Point{})

	// Create an encoder and send some values.
	enc := gob.NewEncoder(&network)
	for i := 1; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
	}

	// Create an decoder and receive some values.
	dec := gob.NewDecoder(&network)
	for i := 1; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Println(result.Hypotenuse())
	}
}
