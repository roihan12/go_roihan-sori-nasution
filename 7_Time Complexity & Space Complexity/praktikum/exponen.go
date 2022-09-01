package main

import (
	"fmt"
	"math"
)

func pow(x,n int) int{
     angka1 := float64(x)
     angka2 := float64(n)
    
    count := math.Pow(angka1, angka2)

    var result int = int(count)

    return result
}

func main(){
    
   fmt.Println(pow(2, 3))  // 8
   fmt.Println(pow(5, 3))  // 125
   fmt.Println(pow(10, 2)) // 100
   fmt.Println(pow(2, 5))  // 32
   fmt.Println(pow(7, 3))  // 343

}