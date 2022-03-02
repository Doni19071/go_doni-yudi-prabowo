package main

import "fmt"

func maxSequence(arr []int) int {
	maxSum := -100000 ^ 100000
	max := 0
	for i := range arr {
		max = max + arr[i]
		if maxSum < max {
			maxSum = max
		}
		if max < 0 {
			max = 0
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(maxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(maxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
