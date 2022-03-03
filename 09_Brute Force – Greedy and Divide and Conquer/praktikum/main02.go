package main

import "fmt"

func coins(money int) []int {
	listMoney := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	hasilTukar := []int{}
	for i := 0; ; {
		if money >= listMoney[i] {
			money = money - listMoney[i]
			hasilTukar = append(hasilTukar, listMoney[i])
			if money == 0 {
				break

			}
		} else {
			i++
		}

	}
	return hasilTukar
}

func main() {
	fmt.Println(coins(123))
	fmt.Println(coins(432))
	fmt.Println(coins(543))
	fmt.Println(coins(7752))
	fmt.Println(coins(15321))
}
