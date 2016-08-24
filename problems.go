package problems

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Problem struct {
	Name      string `json:"name"`
	ShortDesc string `json:"short_desc"`
	Desc      string `json:"long_desc"`
}

func GetList(dirname string) {
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
	json.NewEncoder(os.Stdout).Encode(problems)
}
