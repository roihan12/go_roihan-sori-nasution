package main

import (
	"fmt"
)

func pangkat(base, pangkat int) int {
	var power int = 1;  
	for i := 1; i <= pangkat; i++{      
	power = power * base; 
	} 
	    
	return power
}

func main()  {

	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
	
}