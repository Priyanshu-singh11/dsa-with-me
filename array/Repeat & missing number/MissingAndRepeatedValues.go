package main

import "fmt"

func MissingAndRepeatedValues(nums [][]int) []int {
	n := len(nums)
	var expectedSum, actualSum int = 0, 0
	// finding repeated value
	var repeated, missing int = -1, -1
	set := make(map[int]struct{})
	for _, num := range nums {
		for _, elem := range num {
			actualSum += elem
			if _, ok := set[elem]; ok {
				repeated = elem
			}
			set[elem] = struct{}{}
			// struct{}   {}                  Student    {}
			// ^^^^^^^^   ^^                  ^^^^^^^    ^^
			// Type       Create a value      Type       Create a value
		}
	}

	expectedSum = (n * n) * (n*n + 1) / 2
	missing = expectedSum - actualSum + repeated
	return []int{repeated, missing}
}

func main() {
	nums := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	fmt.Println(MissingAndRepeatedValues(nums))
}
