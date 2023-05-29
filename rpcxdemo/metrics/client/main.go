package main

import (
	"flag"
	"log"
	"math/rand"
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

	go func() {
		for {
			reply := &example.Reply{}
			err := xClient.Call(context.Background(), "Mul", args, reply)
			if err != nil {
				log.Fatalf("failed to call: %v", err)
			}

			//log.Printf("%d * %d = %d", args.A, args.B, reply.C)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}

	}()

	select {}

}
