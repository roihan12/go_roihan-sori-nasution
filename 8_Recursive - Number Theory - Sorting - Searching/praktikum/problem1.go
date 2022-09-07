package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}



func primeX(n int)  int {
	if n == 1 {
		return 2
	}

	if n < 1 {
		return 0
	}

	var countPrime int = 1
	var i int = 3

	for countPrime <= n {
		if checkPrime(i) {
			countPrime++
		}
		i++
	}
	return i - 1
	
}

func checkPrime (number int) bool {
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