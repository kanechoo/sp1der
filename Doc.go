package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"sp1der/models"
)

type Doc struct {
	doc       *goquery.Document
	selectors []models.Selector
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
func (doc *Doc) AddSelectors(selectors ...*models.Selector) *Doc {
	for i := range selectors {
		doc.selectors = append(doc.selectors, *selectors[i])
	}
	return doc
}
func (doc *Doc) AddSelector(selector *models.Selector) *Doc {
	doc.selectors = append(doc.selectors, *selector)
	return doc
}
func (doc *Doc) Result() *[]models.Selector {
	var empty []models.Selector
	if len(doc.selectors) <= 0 {
		return &empty
	}
	for i := range doc.selectors {
		doc.doc.Find((doc.selectors)[i].SelectorQuery).Each(func(j int, selection *goquery.Selection) {
			doc.selectors[i].Text = append(doc.selectors[i].Text, selection.Text())
			if "" != doc.selectors[i].Attr {
				doc.selectors[i].AttrVal = append(doc.selectors[i].AttrVal, selection.AttrOr(doc.selectors[i].Attr, ""))
			}
		})
	}
	return &doc.selectors
}
