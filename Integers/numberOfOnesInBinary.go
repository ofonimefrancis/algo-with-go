package Integers

import "fmt"

/**
	Given a positive integer num, return the sum of its digits.

	Bonus: Can you do it without using strings?

	Example 1
	Input

	num = 123
	Output
	6
**/

func main() {
	fmt.Println(numberOfOnes(15))
}

func numberOfOnes(n int) int {
	count := 0
	for n != 0 {
		quotient := int(n / 2)
		remainder := n % 2

		if remainder == 1 {
			count++
		}

		n = quotient
	}
	return count
}
