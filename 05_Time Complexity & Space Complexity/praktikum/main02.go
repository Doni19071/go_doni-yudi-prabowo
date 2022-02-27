package main

import "fmt"

func pow(x, n int) int {
	output := 1
	for i := 1; i <= n; i++ {
		output *= x
	}
	return output
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
