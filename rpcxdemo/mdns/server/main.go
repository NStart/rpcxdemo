package main

import (
	"flag"
	"log"
	"time"

	"rpcxdemo/mdns/example"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)
	err := s.Register(new(example.Arith), "")
	if err != nil {
		log.Fatalln(err)
	}
	err = s.Serve("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}

}

func addRegistryPlugin(s *server.Server) {
	r := serverplugin.NewMDNSRegisterPlugin("tcp@"+*addr, 8972, metrics.NewRegistry(), time.Minute, "")
	err := r.Start()
	if err != nil {
		log.Fatalln(err)
	}
	s.Plugins.Add(r)
}
