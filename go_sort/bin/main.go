package main

import (
	"fmt"

	bubblesort "github.com/v131v/modern_programming_hw/go_sort/src/bubble_sort"
)

func ask(q string) int {

	fmt.Println(q)

	var n int
	fmt.Scan(&n)

	return n
}

func readAr(n int) []int {

	ar := make([]int, n)

	for i := range ar {
		fmt.Scan(&ar[i])
	}

	return ar
}

func main() {
	n := ask("Enter array size:")

	ar := readAr(n)

	bubblesort.Sort(ar, func(i, j int) bool {
		return ar[i] <= ar[j]
	})

	fmt.Printf("Sorted: %v", ar)
}
