package main

import (
	"flag"
	"log"
	"rpcxdemo/quick/example"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
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
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, map[string]string{"aaa": "from client"})
	ctx = context.WithValue(ctx, share.ResMetaDataKey, make(map[string]string))
	err = xClient.Call(ctx, "Mul", args, reply)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(reply.C)

}
