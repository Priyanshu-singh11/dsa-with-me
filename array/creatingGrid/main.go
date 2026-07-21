package main

import "fmt"

func main() {
	rows, cols := 3, 3
	grid := make([][]int, rows)
	//it creates [ []int *3 ] so it create like that [[] [] []] but it is actually [nil *3] it shows nil becuase it see empty slice
	//fmt.Println(grid[0] == nil) here you can check this
	for i := range grid {
		grid[i] = make([]int, cols)
		// now here i fill all values in row 0 1 and 2 see how and why?
		// it take grid[0] it means that he selected first row which is slice like that []int
		// now we can create integer inside it so we say inside []int create three []int values
	}

	// Now we fill all vlaues here
	value := 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid[i][j] = value
			value++
			fmt.Println(grid)
		}
	}
}

// The important rule is:

// make([][]int, 3) → element type is []int → default value is nil.
// grid := make([][]int, 3)
// grid
// Index 0 -> nil or []int
// Index 1 -> nil or []int
// Index 2 -> nil or []int

// make([]int, 3) → element type is int → default value is 0.
//grid[1] = make([]int, 3)
// it creates Index 0 ---> [0 0 0] or [0,0,0]
