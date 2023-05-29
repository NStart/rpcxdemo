package main

import (
	"flag"
	"log"
	"time"

	"rpcxdemo/consul/example"

	metrics "github.com/rcrowley/go-metrics"
	cserver "github.com/rpcxio/rpcx-consul/serverplugin"
	"github.com/smallnest/rpcx/server"
)

var (
	addr       = flag.String("addr", "127.0.0.1:8972", "server address")
	consulAddr = flag.String("consulAddr", "127.0.0.1:8500", "consul address")
	basePath   = flag.String("base", "/rpcx_test", "prefix path")
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
	r := &cserver.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ConsulServers:  []string{*consulAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatalln(err)
	}
	s.Plugins.Add(r)
}
