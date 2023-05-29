package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"rpcxdemo/tls/example"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

func main() {
	flag.Parse()

	fmt.Println(*addr)

	cert, err := tls.LoadX509KeyPair("../../keys/server.crt", "../../keys/server.key")
	if err != nil {
		log.Fatalf("fail to load cert %s", err)
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

	s := server.NewServer(server.WithTLSConfig(tlsConfig))
	s.Register(new(example.Arith), "")
	err = s.Serve("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}

}
