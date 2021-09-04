package _16_populating_next_right_ptrs

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	Q := list.New()
	Q.PushBack(root)
	for Q.Len() != 0 {
		count := Q.Len()
		for i := 0; i < count; i++ {
			cur := Q.Remove(Q.Front()).(*Node)
			if i < count-1 { // Skip the front
				cur.Next = Q.Front().Value.(*Node) // Point to successor in queue
			}
			if cur.Left != nil {
				Q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				Q.PushBack(cur.Right)
			}
		}
	}
	return root
}
