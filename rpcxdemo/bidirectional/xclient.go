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

	ch := make(chan *protocol.Message)
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*address, "")
	xclient := client.NewBidirectionalXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption, ch)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(reply.C)

	for msg := range ch {
		fmt.Println(msg.Payload)
	}

}
