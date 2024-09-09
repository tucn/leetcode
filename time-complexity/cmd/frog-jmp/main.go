package main

import "fmt"

func main() {
	fmt.Printf("%v\n", Solution(10, 85, 30))
}

func Solution(X int, Y int, D int) int {
	distance := Y - X
	jumps := distance / D
	if distance%D != 0 {
		jumps++
	}
	return jumps
}
