package main

import (
	"log"
	"net/rpc"

	"rpctest"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &rpctest.Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	quotient := new(rpctest.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	if replyCall := <-divCall.Done; replyCall != nil { // will be equal to divCall
		log.Printf("returned: %v", quotient)
		// check errors, print, etc.
	}

}
