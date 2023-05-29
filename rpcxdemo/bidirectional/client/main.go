package main

import (
	"flag"
	"fmt"
	"log"
	"rpcxdemo/quick/example"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"golang.org/x/net/context"
)

var (
	address = flag.String("addr", "127.0.0.1:8972", "")
)

func main() {
	flag.Parse()

	c := client.NewClient(client.DefaultOption)
	err := c.Connect("tcp", *address)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err = c.Call(context.Background(), "Arith", "Mul", args, reply)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(reply.C)

	ch := make(chan *protocol.Message)
	c.RegisterServerMessageChan(ch)

	for msg := range ch {
		fmt.Println(msg.Payload)
	}

}
