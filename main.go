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
	d := distributor.Distributor[distributor.QueryParams]{}
	var max = 0
	d.NextQueryParams(func() (bool, distributor.QueryParams) {
		if max < 10 {
			max++
			return true, distributor.QueryParams{Url: fmt.Sprintf("https://m.weather.com.cn/mweather/101280101.shtml?s=%d", max)}
		}
		return false, distributor.QueryParams{}
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
	executor := task.Executor[distributor.QueryParams, []byte]{}
	executor.HttpClient(
		util.DefaultHttpClient()).
		ParallelSize(10).
		Sync(&wg).
		Source(&channel.UrlChannel).
		Target(&channel.ExecutorResultChannel).Build().S(func(s *distributor.QueryParams) *string {
		return &s.Url
	}).T(func(b *[]byte) *[]byte {
		return b
	}).Run()
	wg.Wait()
}
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
