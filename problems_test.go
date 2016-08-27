package problems

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	data, _ := ioutil.ReadFile("problem-1/README.md")
	problem := &Problem{
		Name: "problem-1",
	}
	Parse(data, problem)
	fmt.Printf("%+v\n", problem)
}

func TestGetList(t *testing.T) {
	f, err := os.Create("autogen_problems.json")
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}
	GetList(".", f)
}
