package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name  []string
	Score []int
}

func (s Student) Average() float64 {
	var average float64

	for _, score := range s.Score {
		average += float64(score)
	}

	average = average / float64(len(s.Score))
	return average
}

func (s Student) Min() (min int, name string) {
	min = 100

	var index int
	for idx, score := range s.Score {
		if score < min {
			min = score
			index = idx
		}

	}

	return s.Score[index], s.Name[index]
}

func (s Student) Max() (max int, name string) {
	max = 0

	var index int
	for idx, score := range s.Score {
		if score > max {
			max = score
			index = idx
		}

	}

	return s.Score[index], s.Name[index]
}

func main() {

	var a = Student{}

	for i := 0; i < 5; i++ {
		var name string

		fmt.Print("Input " + strconv.Itoa(i+1) + " Student’s Name : ")
		fmt.Scan(&name)
		a.Name = append(a.Name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.Score = append(a.Score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
