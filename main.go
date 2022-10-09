package main

import (
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

const (
	maxC = 10
)

var wg = sync.WaitGroup{}

func main() {
	defer timer("main")()
	// start url distribute
	d := distributor.Distributor[string]{}
	var max = 0
	d.NextQueryParams(func() (bool, string) {
		if max < 10 {
			max++
			return true, fmt.Sprintf("https://m.weather.com.cn/mweather/101280101.shtml?s=%d", max)
		}
		return false, ""
	}).Target(&channel.UrlChannel).Run()
	//start to process extract result
	processor := process.Processor[[]byte]{}
	//my processor
	processor.Source(&channel.ExecutorResultChannel).ParallelSize(10).Sync(&wg).
		Processor(func(x *[]byte, t *process.Transport) *process.Transport {
			doc := Doc{}
			result := doc.FromBytes(x).AddSelectors(&models.Title, &models.Footer).Result()
			t.Data = result
			return t
		}).NextProcessor(func(x *process.Transport) *process.Transport {
		fmt.Printf("%v \n", x.Data.(*[]models.Selector))
		return x
	}).Run()
	//start all executor
	executor := task.HttpExecutor{}
	executor.HttpClient(
		util.DefaultHttpClient()).
		ParallelSize(10).
		Sync(&wg).
		Source(&channel.UrlChannel).
		Target(&channel.ExecutorResultChannel).Run()
	wg.Wait()
}
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
