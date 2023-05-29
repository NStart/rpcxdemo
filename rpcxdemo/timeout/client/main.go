package main

import (
	"flag"
	"log"
	"rpcxdemo/quick/example"
	"time"

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
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
	err = xClient.Call(ctx, "Mul", args, reply)
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Request timed out")
		} else {
			log.Fatalln(err)
		}
	}
	cancelFn()

	log.Println(reply.C)

}
