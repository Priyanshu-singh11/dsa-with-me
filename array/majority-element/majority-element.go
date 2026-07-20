package main

import (
	"fmt"
	"sort"
)

var (
	numbers  []int
	numbers2 []int
)

func BruteForceApproach(nums []int) int {
	// you can keep count := 0 then you want to do j:= 0
	for i := 0; i < len(nums); i++ {
		count := 1 //inside because it reset counts
		//{3, 3, 3, 1, 1, 1, 3}
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count > len(nums)/2 {
			return nums[i]
		}
	}
	return -1
}

func OptimizedApproach(nums []int) int {
	sort.Ints(nums)
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			count++
		} else {
			count = 1
		}
		if count > len(nums)/2 {
			return nums[i]
		}
	}
	return -1
}

func MooreVotingAlgorithm(nums []int) int {
	count := 0
	ans := 0
	//   1, 3, 3, 2, 3, 2
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			ans = nums[i]
		}
		if ans == nums[i] {
			count++
		} else {
			count--
		}
	}

	//if value not exist then it verify
	count2 := 0
	for _, num := range nums {
		if num == ans {
			count2++
		}
	}
	if count2 > len(nums)/2 {
		return ans
	}
	return -1
}

func main() {
	numbers := []int{3, 3, 3, 1, 1, 1, 3}
	numbers2 := []int{1, 3, 3, 2, 3, 2}
	fmt.Println(BruteForceApproach(numbers))
	//after sorting arr := []int{1,1,1,1,3,3,3,3}
	fmt.Println(OptimizedApproach(numbers))
	fmt.Println(MooreVotingAlgorithm(numbers2))
	fmt.Println(MooreVotingAlgorithm(numbers))
}
