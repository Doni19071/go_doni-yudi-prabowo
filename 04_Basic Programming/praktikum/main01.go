//Menghitung Luas Permukaan Tabung

package main

import (
	"fmt"
)

func main() {
	var tinggi, jari2 float64
	const pi float64 = 3.14
	fmt.Printf("Masukkan Tinggi : ")
	fmt.Scan(&tinggi)
	fmt.Printf("Masukkan Jari-Jari : ")
	fmt.Scan(&jari2)

	Luas := 2 * pi * jari2 * (jari2 + tinggi)
	fmt.Printf("Luas Permukaan Tabung : ")
	fmt.Println(Luas)
}
