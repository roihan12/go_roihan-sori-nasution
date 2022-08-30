package main

import "fmt"


func PairSum(arr []int, target int) []int {
	
	number:= map[int]int{}
    for i, num := range arr {
        numberPotensial := target - num
        if j, found := number[numberPotensial]; found {
            return []int{j, i}
        }
        number[num] = i
    }
    return []int{}
	
}
func main() {
	fmt.Println(PairSum([]int{1,2,3,4,6}, 6))
	fmt.Println(PairSum([]int{2,5,9,11}, 11))
	fmt.Println(PairSum([]int{1,3,5,7}, 12))
	fmt.Println(PairSum([]int{1,4,6,8}, 10))
	fmt.Println(PairSum([]int{1,5,6,7}, 6))
}