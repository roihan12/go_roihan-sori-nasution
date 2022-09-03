package main

import "fmt"

func main() {
	fmt.Println(Prima(1))
	fmt.Println(Prima(5))
	fmt.Println(Prima(8))
	fmt.Println(Prima(9))
	fmt.Println(Prima(10))
}

func Prima(number int) int {

	slice := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for i, val := range slice {
		if i + 1 == number {
			return val
}
	}
	return -1
}