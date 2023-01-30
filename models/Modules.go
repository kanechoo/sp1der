package models

type Selector struct {
	Name          string
	SelectorQuery string
	Attr          string
	AttrVal       []string
	Text          []string
}
