package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sp1der/channel"
	"sp1der/dist"
	"sp1der/models"
	"sp1der/unit"
	"sp1der/util"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// GetUrlGenerateFunc edit this func impl url generate func for every website
func GetUrlGenerateFunc(website string) func() string {
	var page = 0
	v2exFunc := func() string {
		if page < 40 {
			page++
			return fmt.Sprintf("https://www.v2ex.com/recent?p=%d", page)
		}
		//end func
		return ""
	}
	m := map[string]func() string{
		"v2ex": v2exFunc,
	}
	return m[website]
}
func main() {
	file, err := os.ReadFile("resources/spider.yaml")
	if nil != err {
		panic(err)
	}
	config := models.SpiderConfig{}
	err = yaml.Unmarshal(file, &config)
	if nil != err {
		panic(err)
	}
	for _, website := range config.Websites {
		if true == website.Enabled {
			taskConfig := models.TaskConfig{
				HttpClientPoolSize:    website.Walker.HttpClientPoolSize,
				SleepSecond:           website.Walker.SleepSecond,
				ProcessorPoolSize:     website.Processor.ProcessorPoolSize,
				ProcessorSelectorFile: website.Processor.ProcessorSelectorFile,
				ExportCsvLocation:     website.ExportLocation.Csv,
				Website:               website.WebsiteName,
			}
			startTask(taskConfig, GetUrlGenerateFunc(website.WebsiteName))
		}
	}
}
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
func startTask(taskConfig models.TaskConfig, fun func() string) {
	defer timer(taskConfig.Website)()
	if nil == fun {
		panic("Url generate func must be define")
	}
	walker := dist.HttpWalker{
		SleepTime:          time.Duration(taskConfig.SleepSecond) * time.Second,
		UrlChan:            &channel.UrlChan,
		WaitGroup:          &wg,
		DocumentChan:       &channel.DocumentChan,
		HttpClientPoolSize: taskConfig.HttpClientPoolSize,
	}
	walker.SetUrlGenerateFunc(fun).Walk()
	processor := unit.Processor{
		DocumentChan:     &channel.DocumentChan,
		ParallelSize:     taskConfig.ProcessorPoolSize,
		WaitGroup:        &wg,
		SelectorYamlFile: taskConfig.ProcessorSelectorFile,
	}
	items := processor.ExecuteSelectorQuery()
	wg.Wait()
	//wait all unit done then export csv
	util.CsvExport(items, taskConfig.ExportCsvLocation)
}
