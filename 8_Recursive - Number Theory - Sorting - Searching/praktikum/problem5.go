package main

import "fmt"

func main() {

	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))

	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))

	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))

	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))

}

func FindMinAndMax(arr []int) string {

	MinAndMax:= func (a []int) (min int, max int, indexMin int, indexMax int) {

	min = a[0]
	max = a[0]
	indexMin , indexMax = 0, 0
	for i, value := range a {
		if value < min {
			min = value
			indexMin = i
		}
		if value > max {
			max = value
			indexMax = i
		}
	}
	return min, max, indexMax, indexMin
}

	min, max, indexMax, indexMin := MinAndMax(arr)
	result := fmt.Sprintf("min: %d index: %d max: %d index: %d", min, indexMin, max, indexMax)
	return result
}