package main

import (
	"fmt"
	"strings"
)

func caesarRune(r rune, shift uint) rune {
	s := rune(shift % 26)

	if s == 0 {
		return r
	}

	return (((r - 'a') + s) % 26) + 'a'
}

func caesar(offset uint, input string) string {
	var builder strings.Builder

	for _, r := range input {
		if r >= 'a' && r <= 'z' {
			r = caesarRune(r, offset)
		}

		builder.WriteRune(r)
	}

	return builder.String()
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
