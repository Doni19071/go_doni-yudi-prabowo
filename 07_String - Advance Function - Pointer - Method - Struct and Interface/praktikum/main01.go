package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) string {
	var hasil string
	if len(a) < len(b) {
		str := strings.Contains(b, a)
		if str == true {
			hasil = a
		}
	} else {
		str := strings.Contains(a, b)
		if str == true {
			hasil = b
		}

	}
	return hasil
}

func main() {
	fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGOORO", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("KUPU-KUPU", "KUPU"))
	fmt.Println(compare("ILALANG", "ILA"))
}
