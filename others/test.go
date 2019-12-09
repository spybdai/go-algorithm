package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(GetProducts(numbers))
	numbers = []int{1, 0, 3, 4}
	fmt.Println(GetProducts(numbers))
	numbers = []int{1, 0, 3, 0}
	fmt.Println(GetProducts(numbers))
}

func GetProducts(numbers []int) (products []int) {

	count := len(numbers)
	products = make([]int, count)
	left := 1
	right := 1
	for index, _ := range numbers {
		products[index] = 1
	}

	for index, _ := range numbers {
		products[index] *= left
		left *= numbers[index]
		products[count-1-index] *= right
		right *= numbers[count-1-index]
	}
	return products
}
