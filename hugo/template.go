package hugo

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/yamadatt/movabletype"
)

const pageTpl = `
+++
date = "{{.Date}}"
draft = {{.Draft}}
title = "{{.Title}}"
tags = [{{ StringsJoin .Tags "," }}]
image = "{{ .Image }}"
+++
{{.Content}}
`

type HugoPage struct {
	Date    string
	Draft   bool
	Title   string
	Tags    []string
	Image   string
	Content string
}

func CreateHugoPage(entry *movabletype.Entry) HugoPage {

	tags := make([]string, len(entry.Category))
	for i, s := range entry.Category {
		tags[i] = fmt.Sprintf(`"%s"`, s)
	}

	return HugoPage{
		Date:    entry.Date.String(),
		Draft:   entry.Status != "Publish",
		Title:   entry.Title,
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
