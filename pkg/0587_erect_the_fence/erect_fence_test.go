package _0587_erect_the_fence

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestOuterTrees1(t *testing.T) {
	points := [][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}}
	expected := [][]int{{1, 1}, {2, 0}, {3, 3}, {2, 4}, {4, 2}}
	sort.Slice(expected, func(i, j int) bool {
		return expected[i][0] < expected[j][0] || (expected[i][0] == expected[j][0] && expected[i][1] < expected[j][1])
	})
	result := outerTrees(points)
	require.Equal(t, expected, result)
}
