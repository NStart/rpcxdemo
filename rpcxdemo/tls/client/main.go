package main

import (
	"crypto/tls"
	"flag"
	"log"
	"rpcxdemo/tls/example"

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

	option := client.DefaultOption
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	option.TLSConfig = config
	xClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
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
