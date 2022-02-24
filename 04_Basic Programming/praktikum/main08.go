//Cetak Tabel Perkalian

package main

import "fmt"

func tableX(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			hasil := j * i
			fmt.Printf("%d ", hasil)
		}
		fmt.Printf("\n")
	}
}

func main() {
	tableX(9)
}
