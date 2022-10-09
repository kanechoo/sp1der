package models

var (
	Title = Selector{
		SelectorQuery: "a.topic-link",
		Attr:          "",
		Text:          nil,
		AttrVal:       nil,
	}
	Footer = Selector{
		SelectorQuery: "a.count_livid",
		Attr:          "",
		AttrVal:       nil,
		Text:          nil,
	}
)
