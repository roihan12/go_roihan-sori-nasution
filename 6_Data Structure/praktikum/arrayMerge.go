package main

import "fmt"



func ArrayMerge(arrayA, arrayB []string) []string {
	
	check := make(map[string]int)
	arrayC:= append(arrayA, arrayB...)
	res := make([]string,0)
	for _, val := range arrayC{
		check[val] = 1
	}

	for huruf, _ := range check {
		res = append(res,huruf)
	}

	return res
}
func main()  {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}