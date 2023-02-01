package models

type SelectorQuery struct {
	Name         string `yaml:"name"`
	Home         string `yaml:"home"`
	Description  string `yaml:"description"`
	ItemSelector string `yaml:"item-selector"`
	Selectors    []struct {
		TextSelector string `yaml:"selector"`
		KeyName      string `yaml:"name"`
		AttrSelector string `yaml:"attr"`
		AttrPrefix   string `yaml:"attr-prefix"`
		AttrSuffix   string `yaml:"attr-suffix"`
		TextPrefix   string `yaml:"text-prefix"`
		TextSuffix   string `yaml:"text-suffix"`
	} `yaml:"selectors"`
}
