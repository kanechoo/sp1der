package models

type Selector struct {
	Query   string
	Attr    string
	AttrVal string
	Text    []string
	Indexes []int
}
