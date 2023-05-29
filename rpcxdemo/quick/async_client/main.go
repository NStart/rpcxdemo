package main

import (
	"flag"
	"fmt"
	"log"
	"rpcxdemo/quick/example"

	"github.com/smallnest/rpcx/client"
	"golang.org/x/net/context"
)

var (
	address = flag.String("addr", "127.0.0.1:8972", "")
)

func main() {
	flag.Parse()

	d, err := client.NewPeer2PeerDiscovery("tcp@"+*address, "")
	if err != nil {
		log.Fatalln(err)
	}

	xClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xClient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	call, err := xClient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalln(err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalln(replyCall.Error)
	} else {
		fmt.Println(reply.C)
	}

}
