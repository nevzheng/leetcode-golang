package _617_merge_two_binary_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(x, y *TreeNode) *TreeNode {
	if x != nil && y != nil {
		return &TreeNode{
			x.Val + y.Val,
			mergeTrees(x.Left, y.Left),
			mergeTrees(x.Right, y.Right)}
	}
	if x != nil {
		return x
	}
	return y
}
