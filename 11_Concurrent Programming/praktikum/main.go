package main

import (
	"fmt"
	"time"
)

func getChars(s string) {
	for _, c := range s {
		fmt.Printf("%c ", c)
		time.Sleep(10 * time.Millisecond)
	}
}

func getDigits(digits []int) {
	for _, d := range digits {
		fmt.Printf("%d \n", d)
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	go getChars("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
	go getDigits([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	time.Sleep(1000 * time.Millisecond)
}
