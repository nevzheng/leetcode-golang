package _077_combinations

func dfs(pre []int, idx, n, depth int, ans *[][]int) {
	if depth == 0 {
		cpy := make([]int, len(pre))
		copy(cpy, pre)
		*ans = append(*ans, cpy)
		return
	}
	for i := idx; i <= n; i++ {
		pre = append(pre, i)
		dfs(pre, i+1, n, depth-1, ans)
		pre = pre[:len(pre)-1] // PopBack
	}
}

// combine returns all possible combinations of k numbers in the range 1..n
// To do this we will do DFS traversal on each level we will choose an element
// and then recurse down the tree until we have k elements
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	pre := make([]int, 0, n)
	dfs(pre, 1, n, k, &ans)
	return ans
}
