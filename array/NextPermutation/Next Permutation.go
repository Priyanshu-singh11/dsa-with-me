package main

import "fmt"

// left=0 right=4
// 34519 -> 94513 -> 91543 ->
//91543
func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func nextPermutation(nums []int) {

	var n, pivot int = len(nums), -1

	// 1, 2, 3, 6, 5, 4

	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			pivot = i
			break
		}
	}
	//n-2 because we use numsnums[i+1]

	if pivot == -1 {
		reverse(nums, 0, n-1)
		return
	}
	// 1, 2, 3, 6, 5, 4
	//       ^pivot   ^n-1
	for i := n - 1; i > pivot; i-- {
		if nums[i] > nums[pivot] {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			break
		}
	}

	reverse(nums, pivot+1, n-1)

}

func main() {
	var arr []int = []int{1, 2, 3, 6, 5, 4}
	nextPermutation(arr)
	fmt.Println(arr)
}
