package main

import (
	"fmt"
)

func pairSum(arr []int, target int) []int {
	var x []int
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				x = append(x, i)
				x = append(x, j)

			}
		}
	}
	return x
}

func main() {
	fmt.Println(pairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(pairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(pairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(pairSum([]int{1, 5, 6, 7}, 6))
}
