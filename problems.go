package problems

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type Problem struct {
	Name      string            `json:"name"`
	Title     string            `json:"title"`
	ShortDesc string            `json:"short_desc"`
	FullDesc  string            `json:"full_desc"`
	Templates map[string]string `json:"templates"`
}

func GetList(dirname string, w io.Writer) (map[string]*Problem, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	problems := map[string]*Problem{}
	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), "problem") {
			problem := &Problem{
				Name: file.Name(),
			}
			input, err := ioutil.ReadFile(filepath.Join(dirname, file.Name(), "README.md"))
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			Parse(input, problem)
			problems[problem.Name] = problem
		}
	}

	b, err := json.MarshalIndent(problems, "", "\t")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	w.Write(b)
	return problems, nil
}
