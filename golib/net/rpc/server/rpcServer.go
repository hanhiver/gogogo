package server

import "fmt"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
    fmt.Printf("Call get with A = %d, B = %d\n", args.A, args.B)
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		quo.Quo = 0
		quo.Rem = 0
		return nil
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
