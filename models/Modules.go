package models

type Selector struct {
	Key           string
	SelectorQuery string
	Attr          string
	AttrVal       []string
	Text          []string
}
