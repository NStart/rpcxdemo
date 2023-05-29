package main

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"rpcxdemo/gateway/example"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

func main() {
	flag.Parse()

	fmt.Println(*addr)

	s := server.NewServer()
	s.Register(new(example.Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}

}
