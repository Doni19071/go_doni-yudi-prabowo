package main

import "fmt"

func swap(a, b *int) {
	swapA := *a
	swapB := *b
	*a = swapB
	*b = swapA
}

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}
