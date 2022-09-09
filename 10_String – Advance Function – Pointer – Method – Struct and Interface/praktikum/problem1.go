package main

import (
	"fmt"
)


func Compare(a, b string) string  {

	var kataPanjang, kataPendek, result string
	
	if len(a) < len(b) {
		kataPendek = a
		kataPanjang = b
	} 
	
	if len(a) > len(b) {
		kataPendek = b
		kataPanjang = a
	}

	for i := 0; i < len(kataPanjang); i++ {
		substring := kataPanjang[0:i]
		if substring == kataPendek {
			result = substring
		}
	}
	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGGORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
