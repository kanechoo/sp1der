package models

type SelectorQuery struct {
	ParentSelector string
	ItemSelector   []Selector
}
