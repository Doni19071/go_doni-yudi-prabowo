//Palindrome

package main

import "fmt"

func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(palindrome("civic"))       //true
	fmt.Println(palindrome("katak"))       //true
	fmt.Println(palindrome("kasur rusak")) //true
	fmt.Println(palindrome("mistar"))      //false
	fmt.Println(palindrome("lion"))        //false

}
