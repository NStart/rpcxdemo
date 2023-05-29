package main

import (
	"flag"
	"log"
	"rpcxdemo/quick/example"

	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "127.0.0.1:8972", "server address")
	addr2 = flag.String("addr2", "127.0.0.1:9981", "server address")
)

func main() {
	flag.Parse()

	go createServe(*addr1)
	go createServe(*addr2)

	select {}

}

func createServe(addr string) {
	s := server.NewServer()
	s.Register(new(example.Arith), "")
	err := s.Serve("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
}
