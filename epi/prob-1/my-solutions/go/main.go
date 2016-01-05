package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string) int {
	m := map[rune]int{}
	for _, ch := range s {
		m[ch] += 1
	}
	oddCount := 0
	for _, v := range m {
		if v%2 == 1 {
			oddCount += 1
			if oddCount > 1 {
				return 0
			}
		}
	}
	if oddCount > 1 {
		return 0
	} else {
		return 1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(solve(scanner.Text()))
	}
}
