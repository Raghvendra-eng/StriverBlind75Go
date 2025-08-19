package main

import (
	"fmt"
	"slices"
)

func main() {
	// Create a new instance of the logger
	fmt.Println("Git Repository to solve Strivers Blind 75 problems in go")
	slice := []int{1, 2, 3, 4, 5, 3, 0}
	fmt.Println("Original Slice: ", slice)
	slices.SortFunc(slice, func(a, b int) int {
		return a - b
	})
	fmt.Println("Sorted Slice: ", slice)
}
