package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int  {

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
	
	var t3 = []int{}
	 res := t2[0];
        for i := 1; i < len(t2); i++{
            res = res ^ t2[i];
			t3 = append(t3, res)
}
       
return t3;


	// var t3 = []int{}
	// lenNums := len(t2)
	// res := 0
	// for i := 0; i < lenNums; i++ {
	// 	res = res ^ t2[i]
	// 	t3 = append(t3, res)

	// }

	// return t3


	// pass := []string{}
	// for i := 0; i < len(berubah); i++ {
	// 	for j := i + 1; j < len(berubah); j++ {
	// 		if berubah[i] == berubah[j] {
	// 			berubah[j] = "NULL"
	// 			break
	// 		}
	// 		if j == len(berubah)-1 && berubah[i] != "NULL" {
	// 			pass = append(pass, berubah[i])
	// 		}
	// 		if i == len(berubah)-2 && berubah[j] != "NULL" {
	// 			pass = append(pass, berubah[j])
	// 		}
	// 	}
	// }

	// var t = []int{}
	// for _, i := range pass {
	// 	j, err := strconv.Atoi(i)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	t = append(t, j)
	// }
	// return t
}




func main() {
	fmt.Println(munculSekali("1234123")) 
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}