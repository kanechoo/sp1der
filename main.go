package main

import (
	"fmt"
	"sp1der/channel"
	"sp1der/distributor"
	"sp1der/models"
	"sp1der/process"
	"sp1der/task"
	"sp1der/util"
	"sp1der/v2ex"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	defer timer("main")()
	// start url distribute
	d := distributor.Distributor{}
	var max = 0
	d.NextUrlFunc(func() string {
		if max < 5 {
			max++
			println(fmt.Sprintf("https://www.v2ex.com/recent?p=%d", max))
			return fmt.Sprintf("https://www.v2ex.com/recent?p=%d", max)
		}
		return ""
	}).Target(&channel.HttpUrlChannel).Run()
	//start all executor
	executor := task.HttpExecutor{}
	executor.HttpClient(util.DefaultHttpClient()).
		ParallelExecutorSize(3).
		SleepDuration(2 * time.Second).
		Sync(&wg).UrlChannel(&channel.HttpUrlChannel).
		ResultChannel(&channel.HtmlDocChannel).Run()
	//start to process extract result
	processor := process.Processor{}
	//my processor
	processor.Source(&channel.HtmlDocChannel).ParallelSize(10).Sync(&wg).Run(func(s *[]byte) {
		doc := Doc{}
		result := doc.ToDoc(s).AddSelectorQuery(models.SelectorQuery{
			ParentSelector: "div.box > div.item", ItemSelector: []models.Selector{v2ex.Title, v2ex.CommentCount, v2ex.Author, v2ex.Topic}}).
			Result()
		for _, selectorResult := range *result {
			for _, value := range *selectorResult.Results {
				print(value.Name + ":" + value.Text)
				print("\t")
			}
			print("\n")
		}
	})
	wg.Wait()
}
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
