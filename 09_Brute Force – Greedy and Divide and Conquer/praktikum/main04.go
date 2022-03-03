package main

import "fmt"

func binarySearch(array []int, x int) {
	jumlah := 0
	for i := 0; i < len(array); i++ {
		if x == array[i] {
			fmt.Printf("%d ", i)
			jumlah++
		}
		if array[i] > x {
			break
		}

	}
	if jumlah == 0 {
		fmt.Println(-1)
	}
	fmt.Println()
}

func main() {
	binarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)
	binarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)
	binarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)
	binarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)
}
