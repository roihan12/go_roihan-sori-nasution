package main

import "fmt"

// func printFibonacii(a int, b int, n int) {
// 	var sum int = 0
// 	if n > 0 {
// 		sum = a + b
// 		fmt.Printf("%d ", sum)
// 		a = b
// 		b = sum
// 		printFibonacii(a, b, n-1)
// 	}
// }

// func main() {
// 	var a, b, n int

// 	a = 0
// 	b = 1

// 	fmt.Printf("Enter total number of terms: ")
// 	fmt.Scanf("%d", &n)

// 	fmt.Printf("Fibonacii series is : ")
// 	fmt.Printf("%d\t%d ", a, b)

// 	printFibonacii(a, b, n-2)
// 	fmt.Printf("\n")
// }


func fibonacci(number int) int {
	
	// your code here
		if number == 0 {
			return 0
		} else if number == 1 {
			return 1
		} else {
			return fibonacci(number-1) + fibonacci(number-2)
		}
	}
	

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}