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
	Name      string `json:"name"`
	ShortDesc string `json:"short_desc"`
	Desc      string `json:"long_desc"`
}

func GetList(dirname string, w io.Writer) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
		return
	}
	problems := []*Problem{}
	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), "problem") {
			problem := &Problem{
				Name: file.Name(),
			}
			input, err := ioutil.ReadFile(filepath.Join(file.Name(), "README.md"))
			if err != nil {
				log.Fatal(err)
				return
			}
			lines := strings.Split(string(input), "\n")
			var i, count int
			for i, _ = range lines {
				if strings.HasPrefix(lines[i], "#") {
					count += 1
				}
				if count == 2 {
					break
				}
			}
			problem.ShortDesc = strings.Join(lines[:i], "\n")
			problems = append(problems, problem)
		}
	}
	b, err := json.MarshalIndent(map[string][]*Problem{"problems": problems}, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write(b)
}
