//Faktor Bilangan

package main

import "fmt"

func main() {
	var bilangan int
	fmt.Printf("Masukkan bilangan :")
	fmt.Scanf("%d", &bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}

}
