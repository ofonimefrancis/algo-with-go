package main

import "fmt"

func main() {
	fmt.Println(sqrt(25))
}

func sqrt(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	p := 2

	for p*p <= num {
		p++
	}

	return p - 1
}
