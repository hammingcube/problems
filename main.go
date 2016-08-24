package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Problem struct {
	Name      string
	ShortDesc string
	Desc      string
}

func main() {
	dirname := os.Args[1]
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
	fmt.Printf("%#v\n", problems)
}
