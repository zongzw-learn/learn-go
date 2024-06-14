package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpctest"
)

type Arith int

func (t *Arith) Multiply(args *rpctest.Args, reply *int) error {
	log.Printf("multiply is called: %p", t)
	defer func() { *t++ }()
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *rpctest.Args, quo *rpctest.Quotient) error {
	log.Printf("divide is called: %p", t)
	defer func() { *t++ }()
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func (t *Arith) Count(args *rpctest.Args, r *int) error {
	log.Printf("count is called: %p", t)
	log.Printf("count is: %d", *t)
	*r = int(*t)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	// for range time.Tick(time.Second) {
	// 	log.Printf("t: %d", *arith)
	// }
	select {}
}
