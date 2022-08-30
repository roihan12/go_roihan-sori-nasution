package main

import (
	"fmt"
	"strings"
)



func palindrome(input string) bool {
		// for i := 0; i < len(input); i++ {
		// 	j := len(input)-1-i
		// 	if input[i] != input[j] {
		// 		return false   
		// 	}
		// }
		// return true

		words := strings.Split(input, "")
		var reversed string = ""

		for i := len(words) - 1; i >= 0; i-- {
			reversed += words[i]
		}
		return reversed == input
	}


func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
	
}