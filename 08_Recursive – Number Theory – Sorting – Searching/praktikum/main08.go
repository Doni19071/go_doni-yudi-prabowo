package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func mostAppearItem(items []string) map[int]int {
	ints := make([]int, len(items))
	dict := make(map[int]int)
	for _, num := range ints {
		dict[num] = dict[num] + 1

	}
	return dict
}

func main() {
	fmt.Println(mostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
}
