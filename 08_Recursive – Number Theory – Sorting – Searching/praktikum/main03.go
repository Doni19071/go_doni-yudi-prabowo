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
	i := start + 1
	n := high * wide
	j := 1
	hasil := 0
	for ; ; i++ {
		if cekPrime(i) {
			count++
			hasil += i
			fmt.Printf("%d ", i)
			if count == j*wide {
				fmt.Println("")
				j++
			}
		}
		if count == n {
			break
		}
	}
	fmt.Printf("%d \n\n", hasil)
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}
