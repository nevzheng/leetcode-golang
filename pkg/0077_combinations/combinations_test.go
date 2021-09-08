package _077_combinations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinations1(t *testing.T) {
	n := 4
	k := 2
	result := combine(n, k)
	expected := [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}
	require.Equal(t, expected, result)
}
