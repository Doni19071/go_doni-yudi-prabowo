//Grade Nilai

package main

import "fmt"

func main() {
	var StudentScore int = 88
	fmt.Println(StudentScore)
	switch {
	case StudentScore <= 34:
		fmt.Println("Nilai E")
	case StudentScore <= 49:
		fmt.Println("Nilai D")
	case StudentScore <= 64:
		fmt.Println("Nilai C")
	case StudentScore <= 79:
		fmt.Println("Nilai B")
	case StudentScore <= 100:
		fmt.Println("Nilai A")
	default:
		fmt.Println("Nilai Belum Ada")
	}
}
