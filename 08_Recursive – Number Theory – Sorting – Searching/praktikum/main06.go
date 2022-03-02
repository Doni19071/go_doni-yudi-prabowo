package main

import (
	"fmt"
	"sort"
)

func maxBuy(money int, price []int) {
	sort.Ints(price)
	totalharga := 0
	totalbarang := -1
	for i := 0; i < len(price); i++ {
		if totalharga <= money {
			totalharga += price[i]
			totalbarang++
		} else {
			break
		}
	}

	fmt.Println(totalbarang)
}

func main() {
	maxBuy(50000, []int{25000, 25000, 10000, 14000})
	maxBuy(30000, []int{15000, 10000, 12000, 5000, 3000})
	maxBuy(10000, []int{2000, 3000, 1000, 2000, 10000})
	maxBuy(4000, []int{7500, 3000, 2500, 2000})
	maxBuy(0, []int{10000, 30000})

}
