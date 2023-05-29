package main

import (
	"flag"
	"log"
	"rpcxdemo/consul/example"

	cclient "github.com/rpcxio/rpcx-consul/client"
	"github.com/smallnest/rpcx/client"
	"golang.org/x/net/context"
)

var (
	consulAddr = flag.String("consulAddr", "127.0.0.1:8500", "consul address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	d, err := cclient.NewConsulDiscovery(*basePath, "Arith", []string{*consulAddr}, nil)
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
