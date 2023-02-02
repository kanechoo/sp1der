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
	worker := dist.HttpWalker{
		SleepTime:          1 * time.Second,
		UrlChan:            &channel.UrlChan,
		WaitGroup:          &wg,
		DocumentChan:       &channel.DocumentChan,
		HttpClientPoolSize: 10,
	}
	worker.UrlStoreChan(&channel.UrlChan).UrlGenerateFunc(func() string {
		if page < 2 {
			page++
			return fmt.Sprintf("https://www.v2ex.com/recent?p=%d", page)
		}
		//end func
		return ""
	}).Walk()
	processor := unit.Processor{
		DocumentChan:     &channel.DocumentChan,
		ParallelSize:     10,
		WaitGroup:        &wg,
		SelectorYamlFile: "resources/v2ex.yaml",
	}
	items := processor.ExecuteSelectorQuery()
	wg.Wait()
	//wait all unit done then export csv
	util.CsvExport(items, "/Users/konchoo/Downloads/test.csv")
}
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
