package docs

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sp1der/models"
)

type DocumentParser struct {
	doc                 *goquery.Document
	selectorQueryConfig *models.SelectorQuery
}

func (doc *DocumentParser) ToDoc(b *[]byte) *DocumentParser {
	reader := bytes.NewReader(*b)
	d, err := goquery.NewDocumentFromReader(reader)
	if nil != err {
		panic(err)
	}
	doc.doc = d
	return doc
}
func (doc *DocumentParser) Doc(document *goquery.Document) *DocumentParser {
	doc.doc = document
	return doc
}
func (doc *DocumentParser) SelectorYamlFile(path string) *DocumentParser {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	selectorQueryConfig := models.SelectorQuery{}
	err = yaml.Unmarshal(b, &selectorQueryConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	doc.selectorQueryConfig = &selectorQueryConfig
	return doc
}
func (doc *DocumentParser) ExecuteSelectorQuery() *[]map[string]string {
	items := make([]map[string]string, 0)
	if nil == doc.selectorQueryConfig {
		return &items
	}
	var fragment *goquery.Selection
	if "" == doc.selectorQueryConfig.ItemSelector {
		fragment = doc.doc.Find("html")
	} else {
		fragment = doc.doc.Find(doc.selectorQueryConfig.ItemSelector)
	}
	fragment.Each(func(i int, selection *goquery.Selection) {
		var item = make(map[string]string, 0)
		for index := 0; index < len(doc.selectorQueryConfig.Selectors); index++ {
			var text string
			var attr string
			if "" == doc.selectorQueryConfig.Selectors[index].TextSelector && "" == doc.selectorQueryConfig.Selectors[index].AttrSelector {
				return
			}
			text = selection.Find(doc.selectorQueryConfig.Selectors[index].TextSelector).First().Text()
			text = doc.selectorQueryConfig.Selectors[index].TextPrefix + text + doc.selectorQueryConfig.Selectors[index].TextSuffix
			item[doc.selectorQueryConfig.Selectors[index].KeyName] = text
			if "" != doc.selectorQueryConfig.Selectors[index].AttrSelector {
				attr = selection.Find(doc.selectorQueryConfig.Selectors[index].TextSelector).First().AttrOr(doc.selectorQueryConfig.Selectors[index].AttrSelector, "")
				attr = doc.selectorQueryConfig.Selectors[index].AttrPrefix + attr + doc.selectorQueryConfig.Selectors[index].AttrSuffix
				item[doc.selectorQueryConfig.Selectors[index].KeyName] = attr
			}
		}
		items = append(items, item)
	})
	return &items
}
