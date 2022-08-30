package main

import "fmt"

// Fungsi mengecek bilangan prima
func primeNumber(number int) bool {
	// var bil bool = true;

	// for i := 2; i < number; i++ {
	// 		if number%i == 0 {
	// 			bil = false;
	// 		}
	// 	}            

	// return bil

	if number == 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}

	}

	return true
}


func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))

}