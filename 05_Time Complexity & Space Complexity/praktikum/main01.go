package main

import "fmt"

func primeNumber(number int) bool {

	if number == 1 {
		return false
	}

	if number%2 == 0 {
		return false
	}
	for i := 3; i < number; i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return number > 1
}

func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
