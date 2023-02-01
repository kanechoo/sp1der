package main

import (
	"fmt"
	"sp1der/channel"
	"sp1der/dist"
	"sp1der/unit"
	"sp1der/util"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	defer timer("main")()
	// start url distribute
	var page = 0
	d := dist.Distributor{}
	d.UrlChan(&channel.UrlChan).NextUrlFunc(func() string {
		if page < 1 {
			page++
			return fmt.Sprintf("https://www.v2ex.com/recent?p=%d", page)
		}
		return ""
	}).SyncWaitGroup(&wg).
		HttpClientPoolSize(10).
		RequestSleepTime(2 * time.Second).
		DocumentChan(&channel.DocumentChan).
		Distribute()
	processor := unit.Processor{}
	//my unit
	/*items := make([]map[string]string, 0)
	unit.DocumentChan(&channel.DocumentChan).Parallel(10).SyncWaitGroup(&wg).CallBack(func(s *[]byte) {
		doc := docs.DocumentParser{}
		item := doc.ToDoc(s).SelectorYamlFile("resources/v2ex.yaml").executeSelectorQuery()
		items = append(items, *item...)
	})*/
	items := processor.DocumentChan(&channel.DocumentChan).Parallel(10).
		SyncWaitGroup(&wg).
		SelectorYamlFile("resources/v2ex.yaml").ExecuteSelectorQuery()
	wg.Wait()
	//wait all unit done then export csv
	util.ExportToCsv(items, "/Users/konchoo/Downloads/test.csv", []string{"标题", "评论数", "作者", "话题", "链接"})
}
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
