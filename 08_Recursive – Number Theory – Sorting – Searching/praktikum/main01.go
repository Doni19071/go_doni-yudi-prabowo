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
func primeX(n int) int {
	if n == 0 {
		return 0
	}
	count := 0
	i := 2
	for ; ; i++ {
		if cekPrime(i) {
			count++
		}
		if count == n {
			break
		}
	}
	return i
}
func main() {
	fmt.Println(primeX(2))
}
