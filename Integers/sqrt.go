package Integers

/**
	Given an integer x, find it’s square root.
	Do not use the languages standard libray sqrt()
**/

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

func sqrtBinarySearch(num int) int {
	if num < 2 {
		return num
	}

	start := 1
	end := num
	ans := 0

	for start <= end {
		middle := (start + end) / 2

		if middle*middle == num {
			return middle
		}

		if middle*middle > num {
			end = middle - 1
		}

		if middle*middle < num {
			start = middle + 1
			ans = middle
		}
	}

	return ans
}
