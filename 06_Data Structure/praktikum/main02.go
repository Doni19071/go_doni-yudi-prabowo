package main

import (
	"fmt"
	"strconv"
	"strings"
)

func angka1x(angka string) map[int]int {
	var angka2 []string = strings.Split(angka, "")

	ints := make([]int, len(angka2))

	for i, s := range angka2 {
		ints[i], _ = strconv.Atoi(s)
	}

	dict := make(map[int]int)
	for _, num := range ints {
		dict[num] = dict[num] + 1

	}
	return dict
}
func main() {
	fmt.Println(angka1x("1234123"))
	fmt.Println(angka1x("76523752"))
	fmt.Println(angka1x("12345"))
	fmt.Println(angka1x("1122334455"))
	fmt.Println(angka1x("0872504"))
}
