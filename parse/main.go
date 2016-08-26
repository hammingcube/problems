package main

import (
	"fmt"
	"github.com/golang-commonmark/markdown"
	"io/ioutil"
)

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
}
