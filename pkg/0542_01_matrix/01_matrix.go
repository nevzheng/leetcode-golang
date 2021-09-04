package _542_01_matrix

func Min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func updateMatrix(mat [][]int) [][]int {
	dp := make([][]int, len(mat))
	for i := range mat {
		dp[i] = make([]int, len(mat[0]))
	}
	// Initialize DP Array
	for i, row := range mat {
		for j, elem := range row {
			if elem == 0 {
				dp[i][j] = 0 // Dist from 0 to self is 0
			} else {
				dp[i][j] = 10_007
			}
		}
	}
	// Forward Pass
	for i, row := range mat {
		for j, elem := range row {
			if elem == 1 {
				if i > 0 {
					dp[i][j] = Min(dp[i][j], 1+dp[i-1][j])
				}
				if j > 0 {
					dp[i][j] = Min(dp[i][j], 1+dp[i][j-1])
				}
			}
		}
	}
	// Reverse Pass
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[0]) - 1; j >= 0; j-- {
			if mat[i][j] == 1 {
				if i < (len(mat) - 1) {
					dp[i][j] = Min(dp[i][j], dp[i+1][j]+1)
				}
				if j < (len(mat[0]) - 1) {
					dp[i][j] = Min(dp[i][j], dp[i][j+1]+1)
				}
			}
		}
	}
	return dp
}
