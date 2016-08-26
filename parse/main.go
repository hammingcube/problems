package main

import (
	"fmt"
	"github.com/golang-commonmark/markdown"
	"io/ioutil"
)

func main() {
	md := markdown.New(markdown.XHTMLOutput(true), markdown.Nofollow(true))
	data, _ := ioutil.ReadFile("../problem-1/README.md")

	templates := map[string]string{}

	parsed := md.Parse(data)
	var startRecording bool
	for _, tok := range parsed {
		//fmt.Printf("%s: %#v\n", tok.Tag(), tok)
		if tok, ok := tok.(*markdown.Inline); ok {
			if tok.Content == "Templates" {
				startRecording = true
			}
		}
		if startRecording {
			if _, ok := tok.(*markdown.HeadingOpen); ok {
				startRecording = false
				continue
			}
			if tok, ok := tok.(*markdown.Fence); ok {
				templates[tok.Params] = tok.Content
			}
		}
	}
	fmt.Printf("%#v\n", templates)
}
