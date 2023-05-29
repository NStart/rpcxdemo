package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"rpcxdemo/bidirectional/example"
	"time"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

var clientConn net.Conn
var connected = false

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	clientConn = ctx.Value(server.RemoteConnContextKey).(net.Conn)
	reply.C = args.A * args.B
	connected = true
	return nil
}

func main() {
	flag.Parse()

	fmt.Println(*addr)

	ln, _ := net.Listen("tcp", ":9981")
	go http.Serve(ln, nil)

	s := server.NewServer()
	s.Register(new(Arith), "")
	go func() {
		err := s.Serve("tcp", *addr)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	for !connected {
		time.Sleep(time.Second)
	}
	fmt.Printf("start to send message to %s", clientConn.RemoteAddr().String())

	for {
		if clientConn != nil {
			err := s.SendMessage(clientConn, "test_service_path", "test_service_method", nil, []byte("abcde"))
			if err != nil {
				fmt.Println(err)
				clientConn = nil
			}
		}
	}
}
