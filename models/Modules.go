package models

type Selector struct {
	SelectorQuery string
	Attr          string
	AttrVal       string
	Text          []string
	Indexes       []int
}
