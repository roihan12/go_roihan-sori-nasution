package main

import (
	"fmt"
)


func Caesar( offset int, input string) string {

	offsetKey := byte(offset%26+26) % 26

	var enkripsi []byte

	for _, val := range input {
		hasil := byte(val)
		if 'A' <= val && val <= 'Z' {
			enkripsi = append(enkripsi, 'A'+(hasil-'A'+offsetKey)%26)
		} else if 'a' <= val && val <= 'z' {
			enkripsi = append(enkripsi, 'a'+(hasil-'a'+offsetKey)%26)
		} else {
			enkripsi = append(enkripsi, hasil)
		}
	}
	return string(enkripsi)
}

func main() {
	fmt.Println(Caesar(3, "abc"))
	fmt.Println(Caesar(2, "alta"))
	fmt.Println(Caesar(10, "alterraacademy"))
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(Caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
