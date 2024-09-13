package main

import "fmt"

func main() {
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	result := make([]int, m+n)

	// Simulate the multiplication process
	// For example, 123 * 456
	//    123
	//   *456
	//  ______
	//      738  <- 6 * 123
	//   615[0]   <- 5 * 123
	//492[0][0]    <- 4 * 123
	//  ______
	//  56088

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// using ascii value to get the digit value
			product := (num1[i] - '0') * (num2[j] - '0')
			p1, p2 := i+j, i+j+1
			sum := int(product) + result[p2]

			result[p2] = sum % 10
			result[p1] += sum / 10
		}
	}

	resultStr := ""
	// Simulate the add of multiple numbers
	for _, digit := range result {
		if !(len(resultStr) == 0 && digit == 0) {
			resultStr += string(digit + '0')
		}
	}

	return resultStr
}
