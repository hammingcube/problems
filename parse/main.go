package main

import (
	"fmt"
	"github.com/golang-commonmark/markdown"
	"io/ioutil"
)

func parseTemplates(tokens []markdown.Token, beg, end int) map[string]string {
	templates := map[string]string{}
	for _, tok := range tokens[beg:end] {
		if tok, ok := tok.(*markdown.Fence); ok {
			templates[tok.Params] = tok.Content
		}
	}
	return templates
}

func main() {
	md := markdown.New(markdown.XHTMLOutput(true), markdown.Nofollow(true))
	data, _ := ioutil.ReadFile("../problem-1/README.md")

	//templates := map[string]string{}

	parsed := md.Parse(data)
	m := []struct {
		Title string
		Beg   int
	}{}
	var content string
	for i, tok := range parsed {
		if tok, ok := tok.(*markdown.Inline); ok {
			content = tok.Content
		}
		if _, ok := tok.(*markdown.HeadingClose); ok {
			m = append(m, struct {
				Title string
				Beg   int
			}{content, i})
		}
	}
	fmt.Printf("%+v\n", m)
	var templates map[string]string
	for i, t := range m {
		switch t.Title {
		case "Templates":
			var end int
			if i == len(m)-1 {
				end = len(parsed)
			} else {
				end = m[i+1].Beg
			}
			templates = parseTemplates(parsed, t.Beg, end)
		}
	}
	fmt.Printf("%+v\n", templates)
}
