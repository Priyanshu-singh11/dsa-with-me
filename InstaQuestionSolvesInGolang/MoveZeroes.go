package main

import "fmt"

func moveZeroes(arr []int) []int {
    index := 0
    for i := 0; i < len(arr); i++ {
        if arr[i] != 0 {
            arr[index] = arr[i]
            index++
        }
    }
    for index < len(arr) {
        arr[index] = 0
        index++
    }
    return arr
}

func main() {
    res2 := moveZeroes([]int{0, 1, 0, 2, 3, 0, 4})
    fmt.Println(res2) // [1 2 3 4 0 0 0]
}
