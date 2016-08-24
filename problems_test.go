package problems

import (
	"os"
	"testing"
)

func TestProblems(t *testing.T) {
	f, err := os.Create("autogen_problems.json")
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}
	GetList(".", f)
}
