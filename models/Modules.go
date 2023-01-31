package models

type Selector struct {
	Key           string
	SelectorQuery string
	Attr          string
	AttrVal       []string
	Text          []string
	TextPrefix    string
	TextSuffix    string
	AttrPrefix    string
	AttrSuffix    string
}
