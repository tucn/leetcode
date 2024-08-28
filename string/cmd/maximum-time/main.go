package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maximumTime("2?:?0"))
	fmt.Println(maximumTime("0?:2?"))
	fmt.Println(maximumTime("??:??"))
	fmt.Println(maximumTime("?9:?7"))
}
func maximumTime(time string) string {
	for i := range time {
		if time[i] != '?' {
			continue
		}
		switch i {
		// 1st hour
		case 0:
			if time[i+1] == '?' {
				time = replaceAtIndex(time, '2', i)
				continue
			}
			// cannot be 24 or more
			if time[i+1] > '3' {
				time = replaceAtIndex(time, '1', i)
				continue
			}
			time = replaceAtIndex(time, '2', i)
			continue
		// 1st min
		case 1:
			if time[i-1] == '?' || time[i-1] == '2' {
				time = replaceAtIndex(time, '3', i)
				continue
			}
			time = replaceAtIndex(time, '9', i)
		// 2nd hour
		case 3:
			time = replaceAtIndex(time, '5', i)
		// 2nd min
		case 4:
			time = replaceAtIndex(time, '9', i)
		}
	}
	return time
}

func replaceAtIndex(input string, replacement byte, index int) string {
	return strings.Join([]string{input[:index], string(replacement), input[index+1:]}, "")
}
