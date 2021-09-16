package _042_trapping_rain_water

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRain1(t *testing.T) {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6
	result := trap(heights)
	require.Equal(t, result, expected)
}

func TestRain2(t *testing.T) {
	heights := []int{4, 2, 0, 3, 2, 5}
	expected := 9
	result := trap(heights)
	require.Equal(t, result, expected)
}
