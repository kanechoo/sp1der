package models

var (
	Title = Selector{
		Query:   ".h15 .weaword td",
		Attr:    "",
		Text:    nil,
		AttrVal: "",
		Indexes: []int{0, 1, 2, 3, 4, 5},
	}
	Footer = Selector{
		Query:   ".weatherCardTop ul li.cur span",
		Attr:    "",
		AttrVal: "",
		Text:    nil,
		Indexes: []int{0},
	}
)
