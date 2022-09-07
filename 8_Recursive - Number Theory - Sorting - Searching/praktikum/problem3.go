package main

import (
	"fmt"
	"math"
)

func main() {
   primaSegiEmpat(2, 3, 13)
   primaSegiEmpat(5, 2, 1)
}

func primaSegiEmpat(wide, high, start int) {
	var sum int = 0 

	for i := 0; i < high; i++ {
		for j := 0; j < wide; j++ {
			start = nextPrime(start + 1)
			sum += start
			fmt.Print(start)
			fmt.Print(" ")
		}

		fmt.Println()
	}
	fmt.Println(sum)
	fmt.Println()
}

func nextPrime(start int) int {

	if checkPrime(start) {
		return start
	}

	return nextPrime(start + 1)
}

func checkPrime(number int) bool {
	if number == 1 {
		return false

	}

	if number%2 == 0 {
		return false
	}

	for i := 3; i < int(math.Sqrt(float64(number))); i += 2 {
		if number%i == 0 {
			return false

		}
	}

	return true
}