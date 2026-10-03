
package main

import "fmt"

func threeSum(arr []int) [][]int {
	var newArr [][]int

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {

				fmt.Println(arr[i], arr[j], arr[k])

				sum := arr[i] + arr[j] + arr[k]

				if sum == 0 {
					newArr = append(newArr, []int{
						arr[i],
						arr[j],
						arr[k],
					})
				}
			}
		}
	}

	return newArr
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, 4}

	res := threeSum(arr)

	fmt.Println(res)
}
