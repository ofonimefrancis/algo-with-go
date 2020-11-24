package ArraysAndString

/**
	* Kadence Algorithm
	*
	*		Initialize:
	*			max_so_far = 0
	*			max_ending_here = 0
	*
	*			Loop for each element of the array
	*				(a) max_ending_here = max_ending_here + a[i]
	*				(b) if(max_so_far < max_ending_here)
	*							max_so_far = max_ending_here
	*				(c) if(max_ending_here < 0)
	*							max_ending_here = 0
	*						return max_so_far
	*
**/

// func main() {
// 	fmt.Println(maxSubArray([]int{9, -3, 1, 17}))
// }

func maxSubArray(arr []int) int {
	var cs int = arr[0]
	var gs int = arr[0]

	for i := 1; i < len(arr); i++ {
		cs = max(arr[i], cs+arr[i])

		if cs > gs {
			gs = cs
		}
	}
	return gs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
