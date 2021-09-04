package _001_two_sum

/// 0001

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	result := twoSum(nums, target)
	require.Equal(t, expected, result)
}
