package _16_populating_next_right_ptrs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPopulateNextRight(t *testing.T) {
	seven := &Node{7, nil, nil, nil}
	six := &Node{6, nil, nil, nil}
	five := &Node{5, nil, nil, nil}
	four := &Node{4, nil, nil, nil}
	three := &Node{3, six, seven, nil}
	two := &Node{2, four, five, nil}
	root := &Node{1, two, three, nil}
	root = connect(root)
	require.NotEqual(t, root, nil)
}
