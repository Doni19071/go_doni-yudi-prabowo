//Play With Asterisk

package main

import "fmt"

func PlayWithAsterik(n int) {
	for i := 1; i <= n; i++ {
		for k := n - i; k > 0; k-- {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func main() {
	PlayWithAsterik(5)
}
