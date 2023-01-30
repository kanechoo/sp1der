package models

type Result struct {
	Value *[]SelectorValue
}
type SelectorValue struct {
	Name string
	Text string
	Attr string
}
