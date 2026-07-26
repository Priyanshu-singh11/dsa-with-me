package main

import "fmt"

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) []int {
	var i, j, idx int = m - 1, n - 1, len(nums1) - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[idx] = nums1[i]
			i--
		} else {
			nums1[idx] = nums2[j]
			j--
		}
		idx--
	}

	for j >= 0 {
		nums1[idx] = nums2[j]
		j--
		idx--
	}
	return nums1
}

func main() {
	var nums1, nums2 []int = []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
	//var nums3, nums4 []int = []int{2, 5, 6, 0, 0, 0}, []int{1, 2, 3}
	fmt.Println(MergeSortedArray(nums1, 3, nums2, 3))
	//fmt.Println(MergeSortedArray(nums3, 3, nums4, 3))
}
