package main

import (
	"fmt"
	"sort"
)

func dragonknight(dragonHead, knightHeight []int) {
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	jumlahKnight := len(knightHeight)
	jumlahNaga := len(dragonHead)
	totalNaga := len(dragonHead)
	totalHeight := 0

	if jumlahNaga > jumlahKnight {
		fmt.Println("knight fall")
	} else {
		for i := 0; i < jumlahNaga; i++ {
			for j := 0; j < jumlahKnight; j++ {
				if dragonHead[i] <= knightHeight[j] {
					totalHeight += knightHeight[j]
					totalNaga--
					break
				} else {
					continue
				}
			}
		}
		if totalNaga == 0 {
			fmt.Println(totalHeight)
		} else {
			fmt.Println("knight fall")
		}
	}

}

func main() {
	dragonknight([]int{5, 4}, []int{7, 8, 4})
	dragonknight([]int{5, 10}, []int{5})
	dragonknight([]int{7, 2}, []int{4, 3, 1, 2})
	dragonknight([]int{7, 2}, []int{2, 1, 8, 5})
}
