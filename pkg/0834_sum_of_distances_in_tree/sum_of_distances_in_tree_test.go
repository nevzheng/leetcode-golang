package _834_sum_of_distances_in_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOuterTrees1(t *testing.T) {
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	expected := []int{8, 12, 6, 10, 10, 10}
	result := sumOfDistancesInTree(n, edges)
	require.Equal(t, expected, result)
}
