package main

import (
	"context"
	"flag"
	"log"
	pb "rpcxdemo/protodemo/proto"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

type Arith int

func (a *Arith) Mul(ctx context.Context, args *pb.Args, reply *pb.Reply) error {
	reply.C = args.A + args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	err := s.Register(new(Arith), "")
	if err != nil {
		log.Fatalln(err)
	}
	err = s.Serve("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}

}
