package _095_unique_binary_search_trees_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generate(lo, hi int) []*TreeNode {
	if lo > hi {
		return []*TreeNode{nil}
	}
	subtrees := make([]*TreeNode, 0)
	for i := lo; i <= hi; i++ {
		lefts := generate(lo, i-1)
		rights := generate(i+1, hi)
		for _, l := range lefts {
			for _, r := range rights {
				subtrees = append(subtrees, &TreeNode{i, l, r})
			}
		}
	}
	return subtrees
}

func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}
