package main

import (
	"flag"
	"log"
	"rpcxdemo/roundrobin/example"

	"github.com/smallnest/rpcx/server"
)

var (
	//由于在windows上无法使用reuseport，所以这边add1和add2需要使用不同的端口
	//并且下面启动服务的协议需要为tcp，用于模拟负载均衡
	//如果在linux上则addr1和addr2的端口可以一致
	//然后指定协议未reuseport
	addr1 = flag.String("addr1", "127.0.0.1:8972", "server address")
	addr2 = flag.String("addr2", "127.0.0.1:9981", "server address")
)

func main() {
	flag.Parse()

	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(example.Arith), "weight=7")
		err := s.Serve("tcp", *addr1)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(example.Arith2), "weight=3")
		err := s.Serve("tcp", *addr2)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	select {}

}
