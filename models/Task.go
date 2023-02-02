package models

type TaskConfigure struct {
	Website               string
	HttpClientPoolSize    int
	SleepSecond           int
	ProcessorPoolSize     int
	ProcessorSelectorFile string
	ExportCsvLocation     string
}
