//Frog

package main

import (
	"fmt"
	"math"
)

func frog(jump []int) int {
	i := 0
	hasil := 0
	for i < len(jump) {
		a := math.Abs(float64(jump[i] - jump[i+1]))
		b := math.Abs(float64(jump[i] - jump[i+2]))
		if a < b {
			hasil += int(a)
			i = i + 1
			if i >= len(jump)-1 {
				break
			}
		} else {
			hasil += int(b)
			i = i + 2
			if i >= len(jump)-1 {
				break
			}
		}
	}

	return hasil
}

func main() {
	fmt.Println(frog([]int{10, 30, 40, 20}))
	fmt.Println(frog([]int{30, 10, 60, 10, 60, 50}))
}
