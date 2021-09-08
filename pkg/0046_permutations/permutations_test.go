package _046_permutations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinations1(t *testing.T) {
	nums := []int{1, 2, 3}
	result := permute(nums)
	expected := [][]int{
		{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
	}
	require.Equal(t, expected, result)
}
