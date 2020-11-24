package DynamicProgramming

/**
	Find the sum of the first n numbers

	Objective fuunction
	f(i) is the summ of the first i elements

	Recurrence
	f(n) = f(n-1) + n
**/

func nSum(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i
	}

	return dp[n]
}
