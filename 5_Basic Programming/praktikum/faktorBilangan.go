package main

import "fmt"

func main() {

	//input
	var x int

	fmt.Println("Masukkan bilangan : ")
	fmt.Scanf("%d\n", &x)

	fmt.Println("Faktor dari", x, "adalah:")
    for i := 1; i <= x; i++ {
		if x % i == 0 {
		fmt.Println(i)
		}
	}
        
}