package main

import "fmt"

func minMax(arr []int) {
	min := 1000 ^ 1000
	max := -1000 ^ 1000
	indexMin := 0
	indexMax := 0
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
			indexMin = i
		}
		if arr[i] > max {
			max = arr[i]
			indexMax = i
		}
	}
	fmt.Println("min:", min, "index:", indexMin, "max:", max, "index:", indexMax)
}

func main() {
	minMax([]int{5, 7, 4, -2, -1, 8})
	minMax([]int{2, -5, -4, 22, 7, 7})
	minMax([]int{4, 3, 9, 4, -21, 7})
	minMax([]int{-1, 5, 6, 4, 2, 18})
	minMax([]int{-2, 5, -7, 4, 7, -20})
}
