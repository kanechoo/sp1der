package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"sp1der/models"
)

type Doc struct {
	doc           *goquery.Document
	selectorQuery *models.SelectorQuery
}

func (doc *Doc) ToDoc(b *[]byte) *Doc {
	reader := bytes.NewReader(*b)
	d, err := goquery.NewDocumentFromReader(reader)
	if nil != err {
		panic(err)
	}
	doc.doc = d
	return doc
}
func (doc *Doc) Doc(document *goquery.Document) *Doc {
	doc.doc = document
	return doc
}
func (doc *Doc) AddSelectorQuery(query models.SelectorQuery) *Doc {
	doc.selectorQuery = &query
	return doc
}
func (doc *Doc) ToResult() *[]map[string]string {
	result := make([]map[string]string, 0)
	if nil == doc.selectorQuery {
		return &result
	}
	doc.doc.Find(doc.selectorQuery.ParentSelector).Each(func(i int, selection *goquery.Selection) {
		var mapData = make(map[string]string, 0)
		for j := 0; j < len(doc.selectorQuery.ItemSelector); j++ {
			var value string
			var attr string
			value = selection.Find(doc.selectorQuery.ItemSelector[j].SelectorQuery).First().Text()
			if "" != doc.selectorQuery.ItemSelector[j].Attr {
				attr = selection.Find(doc.selectorQuery.ItemSelector[j].SelectorQuery).First().AttrOr(doc.selectorQuery.ItemSelector[j].Attr, "")
				mapData[doc.selectorQuery.ItemSelector[j].Key] = attr
			}
			mapData[doc.selectorQuery.ItemSelector[j].Key] = value
		}
		result = append(result, mapData)
	})
	return &result
}
