package _129_sum_root_to_leaf_nums

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	dfs(root, 0, &sum)
	return sum
}

func dfs(root *TreeNode, prefix int, sum *int) {
	if root == nil {
		return
	}
	prefix = prefix*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += prefix
		return
	}
	dfs(root.Left, prefix, sum)
	dfs(root.Right, prefix, sum)
}
