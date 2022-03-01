package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type chiper interface {
	Encode() string
	Decode() string
}

func caesarRune(r rune, shift uint) rune {
	s := rune(shift % 26)

	if s == 0 {
		return r
	}

	return (((r - 'a') + s) % 26) + 'a'
}

func (s *student) Encode() string {
	var nameEncode strings.Builder

	for _, r := range s.name {
		if r >= 'a' && r <= 'z' {
			r = caesarRune(r, 25)
		}

		nameEncode.WriteRune(r)
	}

	return nameEncode.String()
}

func (s *student) Decode() string {
	var nameDecode strings.Builder

	for _, r := range s.nameEncode {
		if r >= 'a' && r <= 'z' {
			r = caesarRune(r, 25)
		}

		nameDecode.WriteRune(r)
	}

	return nameDecode.String()

}

func main() {
	var menu int
	var a = student{}
	var c chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu : ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name" + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong Input name menu")
	}
}
