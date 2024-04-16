package hugo

import (
	"bufio"
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yamadatt/movabletype"
)

const pageTpl = `---
date: {{.Date}}
draft: {{.Draft}}
title: "{{.Title}}"
slug: "{{.Slug}}"
tags: [{{ StringsJoin .Tags "," }}]
image: "{{.Image}}"
feature: "{{.Image}}"
thumbnail: "{{.Image}}"
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
	// fmt.Println(entry.Title)
	// fmt.Println(entry.Basename)
	if entry.Basename == "" {
		entry.Basename = entry.Title
	}

	p := bluemonday.StrictPolicy()
	entry.Title = p.Sanitize(entry.Title)
	entry.Basename = p.Sanitize(entry.Basename)

	// titleに特殊文字があるとyamlエラーになるため、全角に書き換える
	entry.Title = strings.Replace(entry.Title, "\"", "”", -1)
	entry.Title = strings.Replace(entry.Title, "\\", "￥", -1)

	// アイキャッチとして画像が入ってない場合は最初のjpgを入れる
	if entry.Image == "" {
		re := regexp.MustCompile(`figure src="(.*?)"`)

		matches := re.FindAllStringSubmatch(entry.Body, -1)
		for i, match := range matches {
			fmt.Printf("Match %d: %s\n", i+1, match[1]) // match[1] にキャプチャグループのマッチ結果が格納される
			if strings.Contains(match[1], "jpg") == true {
				entry.Image = match[1]
				break
			}
		}

		// entry.Image = match

	} else {
		entry.Image = filepath.Base(entry.Image)
		fmt.Println("アイキャッチがもともとあった")
	}

	//debug
	fmt.Println("entry.Image:", entry.Image)

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
