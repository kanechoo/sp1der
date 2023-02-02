package models

type SpiderConfig struct {
	Websites []struct {
		WebsiteName string `yaml:"website-name"`
		Description string `yaml:"description"`
		Enabled     bool   `yaml:"enabled"`
		Walker      struct {
			SleepSecond        int `yaml:"sleep-second"`
			HttpClientPoolSize int `yaml:"http-client-pool-size"`
		} `yaml:"walker"`
		Processor struct {
			ProcessorPoolSize     int    `yaml:"processor-pool-size"`
			ProcessorSelectorFile string `yaml:"processor-selector-file"`
		} `yaml:"processor"`
		ExportLocation struct {
			Csv string `yaml:"csv"`
		} `yaml:"export-location"`
	} `yaml:"website-list"`
}
