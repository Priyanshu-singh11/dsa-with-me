package main

import "fmt"

func findDuplicate(nums []int) int {
	duplicateValue := 0
	seen := make(map[int]struct{})
	for _, num := range nums {
		_, ok := seen[num]
		if ok {
			duplicateValue = num
			break
		}
		fmt.Println("running with ", num)
		seen[num] = struct{}{}
	}

	return duplicateValue
}

func slowFastPointerMethod(nums []int) int {
	//var ans int = 0
	var slow, fast int = nums[0], nums[0]
	for {
		slow = nums[slow]       //+1
		fast = nums[nums[fast]] //+2
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow] //+1
		fast = nums[fast] //+1
	}
	return slow
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	nums1 := []int{3, 1, 3, 4, 2}
	nums2 := []int{3, 3, 3, 3, 3}
	//fmt.Println(findDuplicate(nums))
	//fmt.Println(findDuplicate(nums1))
	//fmt.Println(findDuplicate(nums2))
	fmt.Println(slowFastPointerMethod(nums))
	fmt.Println(slowFastPointerMethod(nums1))
	fmt.Println(slowFastPointerMethod(nums2))
}
