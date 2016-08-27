package problems

import (
	"github.com/golang-commonmark/markdown"
	"strings"
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

func parseDesc(tokens []markdown.Token, beg, end int) string {
	content := []string{}
	for _, tok := range tokens[beg+1 : end] {
		if tok, ok := tok.(*markdown.Inline); ok {
			content = append(content, tok.Content)
		}
	}
	return strings.Join(content, "\n")
}

func Parse(data []byte, problem *Problem) {
	md := markdown.New(markdown.XHTMLOutput(true), markdown.Nofollow(true))
	//data, _ := ioutil.ReadFile("../problem-1/README.md")

	var templates map[string]string
	var title, shortDesc, fullDesc string

	lines := strings.Split(string(data), "\n")
	var index int
	for index, _ = range lines {
		if strings.HasPrefix(lines[index], "### Templates") {
			break
		}
	}
	fullDesc = strings.Join(lines[:index], "\n")

	// parse rest of the content
	parsed := md.Parse(data)
	type Unit struct {
		Title string
		Beg   int
	}
	m := []Unit{}

	var content string
	for i, tok := range parsed {
		if tok, ok := tok.(*markdown.Inline); ok {
			content = tok.Content
		}
		if _, ok := tok.(*markdown.HeadingClose); ok {
			if tok.Tag() == "h3" {
				m = append(m, struct {
					Title string
					Beg   int
				}{content, i - 1})
			}
		}
	}
	//fmt.Printf("%+v\n", m)
	getBegEnd := func(parsed []markdown.Token, m []Unit, i int) (int, int) {
		var end int
		if i == len(m)-1 {
			end = len(parsed)
		} else {
			end = m[i+1].Beg
		}
		beg := m[i].Beg
		return beg, end
	}

	for i, t := range m {
		switch t.Title {
		case "Templates":
			beg, end := getBegEnd(parsed, m, i)
			templates = parseTemplates(parsed, beg, end)
		case "Short Description":
			beg, end := getBegEnd(parsed, m, i)
			shortDesc = parseDesc(parsed, beg, end)
		}
	}
	title = m[0].Title
	problem.Title = title
	problem.ShortDesc = shortDesc
	problem.FullDesc = fullDesc
	problem.Templates = templates
	//fmt.Printf("%q\n%q\n%q\n%q\n", title, shortDesc, fullDesc, templates)
}
