package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, ": ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Variadic ...
func Variadic() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	fmt.Println("nums: ", nums)
	sum(nums...)
}