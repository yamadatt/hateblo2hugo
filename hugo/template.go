package hugo

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/yamadatt/movabletype"
)

const pageTpl = `---
date: {{.Date}}
draft: {{.Draft}}
title: "{{.Title}}"
slug: "{{.Slug}}"
tags: [{{ StringsJoin .Tags "," }}]
image: "{{.Image}}"
---
{{.Content}}
`

type HugoPage struct {
	Date    string
	Draft   bool
	Title   string
	Slug    string
	Tags    []string
	Image   string
	Content string
}

func CreateHugoPage(entry *movabletype.Entry) HugoPage {

	tags := make([]string, len(entry.Category))
	for i, s := range entry.Category {
		tags[i] = fmt.Sprintf(`"%s"`, s)
	}

	du, _ := time.ParseDuration("-9h")
	d := entry.Date.Add(du)
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	dJST := d.In(jst)
	fmt.Println(entry.Title)
	fmt.Println(entry.Basename)
	if entry.Basename == "" {
		entry.Basename = entry.Title
	}

	// titleに特殊文字があるとyamlエラーになるため、全角に書き換える
	entry.Title = strings.Replace(entry.Title, "\"", "”", -1)
	entry.Title = strings.Replace(entry.Title, "\\", "￥", -1)

	return HugoPage{
		Date:    dJST.Format(time.RFC3339),
		Draft:   entry.Status != "Publish",
		Title:   entry.Title,
		Slug:    entry.Basename,
		Tags:    tags,
		Image:   entry.Image,
		Content: entry.Body,
	}
}

func (p *HugoPage) Render() ([]byte, error) {

	// strings.LastIndex
	tpl := template.Must(template.New("hugo").Funcs(template.FuncMap{"StringsJoin": strings.Join}).Parse(pageTpl))

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	if err := tpl.Execute(writer, *p); err != nil {
		return nil, err
	}
	if err := writer.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
