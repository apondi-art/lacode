package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 5
	fmt.Println(Draw(n))
}

func Draw(n int) []string {
	var result []string
	width := 2*n - 1 // Maximum width of the bottom floor

	for i := 1; i <= n; i++ {
		stars := 2*i - 1
		spaces := (width - stars) / 2
		floor := strings.Repeat(" ", spaces) + strings.Repeat("*", stars) + strings.Repeat(" ", spaces)
		result = append(result, floor)
	}

	return result
}
