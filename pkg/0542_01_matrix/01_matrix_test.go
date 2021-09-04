package _542_01_matrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test01Matrix1(t *testing.T) {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	expected := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	result := updateMatrix(mat)
	require.Equal(t, expected, result)
}
