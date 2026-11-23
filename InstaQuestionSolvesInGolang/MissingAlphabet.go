package main

import "fmt"

func missinAlphabet(str string) string {
	for i := 0; i < len(str)-1; i++ {
		fmt.Println(string(str[i+1]))
		curr := int(str[i])           
		next := int(str[i+1])      
		if next-curr > 1 {
			return string(rune(curr + 1)) 
		}
	}
	return "Not any missing "
}

func main() {
	// res := missinAlphabet("abde")  // returns "c"
	// res1 := missinAlphabet("abcdf") // returns "e"
	res2 := missinAlphabet("abcd")
	fmt.Println(res2) 
}
