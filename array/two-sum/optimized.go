package main

// import "fmt"

// func twoSumDecOrder(arr []int, target int) []int {
// 	left := 0
// 	right := len(arr) - 1
// 	for left < right {
// 		sum := arr[left] + arr[right]
// 		if sum > target {
// 			left++
// 		}
// 		if sum < target {
// 			right--
// 		}
// 		if sum == target {
// 			return []int{arr[left], arr[right]}
// 		}
// 	}
// 	return nil
// }

// func twosumIncOrder(arr []int, target int) []int {
// 	left := 0
// 	right := len(arr) - 1
// 	for left < right {
// 		sum := arr[left] + arr[right]
// 		if sum > target {
// 			right--
// 		}
// 		if sum < target {
// 			left++
// 		}
// 		if sum == target {
// 			return []int{arr[left], arr[right]}
// 		}
// 	}
// 	return nil
// }

// func main() {

// 	//Increasing order
// 	arr := []int{2, 7, 11, 15}
// 	target := 26
// 	result := twosumIncOrder(arr, target)
// 	fmt.Println(result)

// 	//Decreasing order array
// 	arr2 := []int{15, 11, 7, 2}
// 	target2 := 26
// 	result2 := twoSumDecOrder(arr2, target2)
// 	fmt.Println(result2)
// }
