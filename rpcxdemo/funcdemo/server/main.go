package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

func main() {
	flag.Parse()

	fmt.Println(*addr)

	s := server.NewServer()
	s.RegisterFunction("a.fake.service", Mul, "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}

}
