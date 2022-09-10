package main

import (
	"fmt"
	"strings"
)


type mapFrekuensi map[string]int

// Frekuensi menghitung frekuensi setiap string yang diberikan.
func Frekuensi(s string) mapFrekuensi {
	mapping := mapFrekuensi{}
	for _, r := range s {
		mapping[string(r)]++
	}
	return mapping
}
// HitungFrekuensi menghitung frekuensi string secara bersamaan.
func HitungFrekuensi(texts string) {
	result := mapFrekuensi{}
	channel := make(chan mapFrekuensi, len(texts))


	Split := strings.Split(texts, ",") 

	for _, text := range Split {
		go func(text string) {
			channel <- Frekuensi(text)
		}(text)
	}


	for range Split {
		for r, v := range <-channel {
			result[r] += v
		}
	}

	for i, v := range result {
		fmt.Println(i, ":", v)
	}
}

func main() {
	// kode disini
	input := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	HitungFrekuensi(input)
}