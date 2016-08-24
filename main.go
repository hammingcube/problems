package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dirname := os.Args[1]
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), "problem") {
			input, err := ioutil.ReadFile(filepath.Join(file.Name(), "README.md"))
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("%q\n", input)
		}
	}
}
