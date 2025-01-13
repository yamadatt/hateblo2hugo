package transformer

import (
	"github.com/PuerkitoBio/goquery"
)

type JavascriptTransformer struct {
	doc *goquery.Document
}

func (t *JavascriptTransformer) Transform() error {
	t.doc.Find("noscript").Remove()

	return nil
}

// func (s *Selection) Remove() *Selection
// func (s *Selection) Find(selector string) *Selection
