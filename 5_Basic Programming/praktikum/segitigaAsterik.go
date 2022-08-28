package main

import "fmt"

func playWithAsterik(n int) {

	var k int
	for i := 1; i <= n; i++ {
		k = 0
		for space := 1; space <= n-i; space++ {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
			k++
		}
		fmt.Println("")
	}

}

func main() {
	playWithAsterik(5)

}