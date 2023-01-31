package util

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ExportToCSVV2(data *[][]map[string]string, path string, headers []string) {
	list := make([]map[string]string, 0)
	for _, item := range *data {
		list = append(list, item...)
	}
	ExportToCSV(&list, path, headers)
}
func ExportToCSV(data *[]map[string]string, path string, headers []string) {
	//get titles names
	if nil == data || len(*data) <= 0 {
		return
	}
	titles := make([]string, 0)
	records := make([][]string, 0)
	if nil != headers && len(headers) > 0 {
		titles = headers
	} else {
		for key := range (*data)[0] {
			titles = append(titles, key)
		}
	}
	for _, mapData := range *data {
		record := make([]string, 0)
		for i := 0; i < len(titles); i++ {
			record = append(record, mapData[titles[i]])
		}
		records = append(records, record)
	}
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Create file error : %s", err.Error())
		return
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write(titles)
	writer.WriteAll(records)
}
