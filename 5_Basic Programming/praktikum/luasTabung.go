package main

import "fmt"

func main() {
	//input
	var T float64
	var r float64
	const PI float64 = 3.14
	
	fmt.Println("Masukkan tinggi tabung : ")
	fmt.Scanf("%f\n", &T)
	fmt.Println("Masukkan jari jari : ")
	fmt.Scanf("%f\n", &r)

	luasPermukaan := (2.0 * PI * (r * r)) + (2.0 * PI * r * T)
	fmt.Println("Luas Permukaan tabung adalah : ", luasPermukaan)

}