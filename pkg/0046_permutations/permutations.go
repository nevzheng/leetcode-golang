package _046_permutations

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	dfs(nums, 0, &ans)
	return ans
}

func dfs(nums []int, idx int, ans *[][]int) {
	if idx == len(nums) {
		cpy := make([]int, len(nums))
		copy(cpy, nums)
		*ans = append(*ans, cpy)
		return
	}
	for i := idx; i < len(nums); i++ {
		nums[i], nums[idx] = nums[idx], nums[i] // Swap i and idx
		dfs(nums, idx+1, ans)
		nums[i], nums[idx] = nums[idx], nums[i] // Swap Back
	}
}
