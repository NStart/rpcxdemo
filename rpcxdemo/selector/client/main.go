package main

import (
	"flag"
	"log"
	"rpcxdemo/roundrobin/example"
	"sort"
	"strings"
	"time"

	"github.com/smallnest/rpcx/client"
	"golang.org/x/net/context"
)

var (
	address1 = flag.String("addr1", "127.0.0.1:8972", "")
	address2 = flag.String("addr2", "127.0.0.1:9981", "")
)

func main() {
	flag.Parse()

	d, err := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *address1},
		{Key: *address2}})
	if err != nil {
		log.Fatalln(err)
	}

	xClient := client.NewXClient("Arith", client.Failtry, client.ConsistentHash, d, client.DefaultOption)
	defer xClient.Close()
	selector := &alwaysFirstSelector{}
	xClient.SetSelector(selector)

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &example.Reply{}
		err = xClient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(reply.C)
		time.Sleep(1e9)
	}

}

type alwaysFirstSelector struct {
	servers []string
}

// Select 根据路由策略选择要发送请求的节点
func (s *alwaysFirstSelector) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	var ss = s.servers
	if len(ss) == 0 {
		return ""
	}
	return ss[0]
}

func (s *alwaysFirstSelector) UpdateServer(servers map[string]string) {
	var ss = make([]string, 0, len(servers))
	for k := range servers {
		ss = append(ss, k)
	}

	sort.Slice(ss, func(i, j int) bool {
		return strings.Compare(ss[i], ss[j]) < 0
	})
	s.servers = ss
}
