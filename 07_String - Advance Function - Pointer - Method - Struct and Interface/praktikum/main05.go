package main

import "fmt"

type student struct {
	name  []string
	score []int
}

func (s student) Average() float64 {
	result := 0.0
	j := len(s.score)
	for i := 0; i < j; i++ {
		result = result + float64(s.score[i])
	}
	result = result / float64(j)
	return result
}

func (s student) Min() (min int, name string) {
	min = 10 ^ 100000000
	j := len(s.score)
	for i := 0; i < j; i++ {
		if min > s.score[i] {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return min, name
}

func (s student) Max() (max int, name string) {
	max = -10 ^ 100000000
	j := len(s.score)
	for i := 0; i < j; i++ {
		if max < s.score[i] {
			max = s.score[i]
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	var a = student{}

	for i := 1; i < 6; i++ {
		var name string
		fmt.Print("Input " + string(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is : ", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
