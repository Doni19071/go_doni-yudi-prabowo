package main

import (
	"fmt"
	"math"
)

func simpleEquation(a, b, c int) {
	var x, y, z = 0, 0, 0
	fungsi := (a*a - c)
	max := int(math.Sqrt(float64(c)))
	hasil := false
	for i := 1; i <= max; i++ {
		x = i
		for j := 1; j <= max; j++ {
			y = j
			for k := 1; k <= max; k++ {
				z = k
				if 2*(x*y+x*z+y*z) == fungsi {
					fmt.Println(x, y, z)
					hasil = true
				} else {
					z = 0
				}
			}
			if hasil == true {
				break
			}
		}
		if hasil == true {
			break
		}

	}
	if x == 0 || y == 0 || z == 0 {
		fmt.Println("no solution")
	}
}

func main() {
	simpleEquation(1, 2, 3)
	simpleEquation(6, 6, 14)
}
