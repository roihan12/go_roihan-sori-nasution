package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) int  {

	berubah := []string{}
	for i := range angka {
		kata := string(angka[i])
		berubah = append(berubah, kata)
	}

	var t2 = []int{}
	for _, i := range berubah {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	
}




func main() {
	fmt.Println(munculSekali("1234123")) 
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}