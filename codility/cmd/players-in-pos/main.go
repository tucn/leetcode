package main

import "fmt"

func main() {
	fmt.Printf("%v\n", countMoves("<<>>"))
	fmt.Printf("%v\n", countMoves(">>>>"))
}

func countMoves(N string) int {
	maxLength := len(N)
	occupied := make([]bool, maxLength)

	// Make all player true
	for i := range maxLength {
		occupied[i] = true
	}

	const (
		LEFT  = '<'
		RIGHT = '>'
		UP    = '^'
		DOWN  = 'v'
	)

	count := 0

	for i := range maxLength {
		switch N[i] {
		case LEFT:
			if i == 0 {
				count++
				occupied[i] = false
				continue
			}
			if i-1 >= 0 && !occupied[i-1] {
				count++
				occupied[i] = false
				occupied[i-1] = true
			}
		case RIGHT:
			if i == maxLength-1 {
				count++
				occupied[i] = false
				continue
			}
			if i+1 < maxLength && !occupied[i+1] {
				count++
				occupied[i+1] = true
				occupied[i] = false
			}
		case UP, DOWN:
			count++
			occupied[i] = false
		}
	}
	return count

}
