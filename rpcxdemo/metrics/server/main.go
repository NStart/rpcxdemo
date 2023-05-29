package main

import (
	"flag"
	"log"
	"time"

	example "rpcxdemo/metrics/example"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	p := serverplugin.NewMetricsPlugin(metrics.DefaultRegistry)
	s.Plugins.Add(p)
	startMetrics()

	s.Register(new(example.Arith), "")
	s.Serve("tcp", *addr)
}

func startMetrics() {
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry)
	go metrics.CaptureRuntimeMemStats(metrics.DefaultRegistry, time.Second)

	//定期将数据打印到控制台
	go func() {
		for range time.Tick(time.Second * 10) {
			metrics.DefaultRegistry.Each(func(s string, i interface{}) {
				log.Printf("Metric: %s, value %v", s, i)
			})
		}
	}()

}
