package main

import (
	"fmt"
	"math"
)

func cekPrime(n int) bool {
	if n < 2 {
		return false
	}
	angka := true
	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			angka = false
			break
		}
	}
	return angka
}

func primaSegiEmpat(high, wide, start int) {
	if high == 0 || wide == 0 {
		fmt.Println("Tidak ada")
	}
	count := 0
	i := start
	n := high * wide
	for ; ; i++ {
		for j := 0; j < high; j++ {
			for k := 0; k < wide; k++ {
				if cekPrime(i) {
					count++
					fmt.Printf("%d ", i)
				}

			}
			fmt.Printf("\n")
		}
		if count == n {
			break
		}

	}
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}
