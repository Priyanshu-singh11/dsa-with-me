package main

import "fmt"

func hashMap(nums []int, target int) []int {
	//{3,2,4,5,6,7}
	indexMap := make(map[int]int)
	for i, num := range nums {
		sec := target - num
		if index, ok := indexMap[sec]; ok {
			return []int{i, index}
		}
		indexMap[num] = i
	}
	return []int{}
}

func main() {
	num := []int{3, 2, 4, 5, 6, 7}
	target := 9
	fmt.Println(hashMap(num, target))
}
