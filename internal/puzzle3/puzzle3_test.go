package puzzle3_test

import (
	"strconv"
	"testing"

	"aoc2025/internal/puzzle3"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	type testCase struct {
		banks    [][]int
		size     int
		expected int
	}

	testCases := []testCase{
		{
			banks: [][]int{
				{9, 8, 7, 6, 5, 4, 3, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			size:     5,
			expected: 98789,
		},
		{
			banks: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8},
			},
			size:     5,
			expected: 45678,
		},
		{
			banks: [][]int{
				{9, 9, 8, 8, 7, 6},
			},
			size:     4,
			expected: 9988,
		},
		{
			banks: [][]int{
				{9, 8, 7, 6, 5, 4, 3, 9},
			},
			size:     5,
			expected: 98769,
		},
		{
			banks: [][]int{
				{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
				{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
				{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
				{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			},
			size:     12,
			expected: 3121910778619,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, puzzle3.Run(tc.banks, tc.size))
		})
	}
}
