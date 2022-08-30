package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int  {

	angkaString := []string{}
	for i := range angka {
		angka:= string(angka[i])
		angkaString = append(angkaString, angka)
	}

	checkAngka := []string{}
	for i := 0; i < len(angkaString); i++ {
		for j := i + 1; j < len(angkaString); j++ {
			if angkaString[i] == angkaString[j] {
				angkaString[j] = "NULL"
				break
			}
			if j == len(angkaString)-1 && angkaString[i] != "NULL" {
				checkAngka = append(checkAngka, angkaString[i])
			}
			if i == len(angkaString)-2 && angkaString[j] != "NULL" {
				checkAngka = append(checkAngka, angkaString[j])
			}
		}
	}

	var t = []int{}
	for _, i := range checkAngka {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t = append(t, j)
	}
	return t

}




func main() {
	fmt.Println(munculSekali("1234123")) 
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}