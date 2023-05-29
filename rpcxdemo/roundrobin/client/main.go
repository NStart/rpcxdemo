package main

import (
	"flag"
	"log"
	"rpcxdemo/roundrobin/example"
	"time"

	"github.com/smallnest/rpcx/client"
	"golang.org/x/net/context"
)

var (
	address1 = flag.String("addr1", "127.0.0.1:8972", "")
	address2 = flag.String("addr2", "127.0.0.1:9981", "")
)

func main() {
	flag.Parse()

	d, err := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *address1}, {Key: *address2}})
	if err != nil {
		log.Fatalln(err)
	}

	xClient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	defer xClient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &example.Reply{}
		err = xClient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(reply.C)
		time.Sleep(1e9)
	}

}
