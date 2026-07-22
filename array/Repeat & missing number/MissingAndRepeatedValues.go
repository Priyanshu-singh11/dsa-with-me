package main

import "fmt"

func MissingAndRepeatedValues(nums [][]int) []int {
	n := len(nums)
	seen := make(map[int]struct{})
	expectedSum, actualSum := 0, 0
	repeated, missing := 0, 0
	//finding repeated
	for _, row := range nums {
		for _, elem := range row {
			actualSum += elem
			_, ok := seen[elem]
			if ok {
				repeated = elem
			}
			seen[elem] = struct{}{}
		}
	}

	expectedSum = (n * n) * (n*n + 1) / 2
	missing = expectedSum - actualSum + repeated

	//finding missing
	return []int{repeated, missing}
}

func main() {
	nums := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	fmt.Println(MissingAndRepeatedValues(nums))
}
