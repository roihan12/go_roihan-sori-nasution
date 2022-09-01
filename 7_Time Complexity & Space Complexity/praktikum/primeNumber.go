package main

import (
	"fmt"
	"math"
)



func main() {
	fmt.Println(primeNumber(1000000007)) 
	fmt.Println(primeNumber(13))        
	fmt.Println(primeNumber(17))        
	fmt.Println(primeNumber(20))         
	fmt.Println(primeNumber(35))                 
 
}

func primeNumber(number int) bool {

		for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
			if number%i == 0 {
				return false
			}
		}
		return number > 1

	}

