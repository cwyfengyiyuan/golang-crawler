package main

import (
	"fmt"
	"learn/crawler_distributed/config"
	"learn/crawler_distributed/rpcsupport"
	"learn/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T)  {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second*6)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url:    "https://album.zhenai.com/u/1373315316",
		Parser: worker.SerializedParser{
			Name:config.ParseProfile,
			Args:[]interface{}{"句号","女士"},
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	}else {
		fmt.Println(result)
	}
}
