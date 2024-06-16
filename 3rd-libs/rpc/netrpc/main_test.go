package main

import (
	"log"
	"net/rpc"
	"testing"
)

func TestGORPC(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	<-divCall.Done // will be equal to divCall
	log.Printf("returned: %v", quotient)
	// check errors, print, etc.

	c := 0
	countCall := client.Go("Arith.Count", args, &c, nil)
	if countReply := <-countCall.Done; countReply.Error != nil {
		log.Println(countReply.Error.Error())
	} else {
		log.Printf("called %d", c)
	}

	err = client.Call("Arith.Count", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("called %d times", reply)
}
