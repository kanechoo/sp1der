package models

type SelectorResult struct {
	Results *[]SelectorValue
}
type SelectorValue struct {
	Name string
	Text string
	Attr string
}
