package main

import (
	"encoding/json"
	"fmt"
	"sp1der/channel"
	"sp1der/distributor"
	"sp1der/models"
	"sp1der/process"
	"sp1der/task"
	"sp1der/util"
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
		if max < 1 {
			max++
			println(fmt.Sprintf("https://www.v2ex.com/recent?p=%d", max))
			return fmt.Sprintf("https://www.v2ex.com/recent?p=%d", max)
		}
		return ""
	}).Target(&channel.HttpUrl).Run()
	//start all executor
	executor := task.HttpExecutor{}
	executor.HttpClient(util.DefaultHttpClient()).ParallelSize(10).Sync(&wg).Url(&channel.HttpUrl).FetchToDoc(&channel.HtmlDoc).Run()
	//start to process extract result
	processor := process.Processor{}
	//my processor
	processor.Source(&channel.HtmlDoc).ParallelSize(10).Sync(&wg).Run(func(s *[]byte) {
		doc := Doc{}
		result := doc.ToDoc(s).AddSelectors(&models.Title, &models.Footer).Result()
		marshal, _ := json.Marshal(result)
		fmt.Printf("%s\v", string(marshal))
	})
	wg.Wait()
}
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
