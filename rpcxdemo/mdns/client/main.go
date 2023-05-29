package main

import (
	"flag"
	"log"
	"rpcxdemo/mdns/example"
	"time"

	"github.com/smallnest/rpcx/client"
	"golang.org/x/net/context"
)

var (
	basePath = flag.String("base", "/rpcx_test/Arith", "prefix path")
)

func main() {
	flag.Parse()

	d, err := client.NewMDNSDiscovery(*basePath, 10*time.Second, 10*time.Second, "")
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

	log.Println(reply.C)

}
