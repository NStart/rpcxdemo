package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"rpcxdemo/inprocess/example"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	err := s.Register(new(example.Arith), "")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		err = s.Serve("tcp", *addr)
		if err != nil {
			log.Fatalln(err)
		}

	}()

	d, err := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
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
	err = xClient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(reply.C)

}
