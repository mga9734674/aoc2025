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
		expected int
	}

	testCases := []testCase{
		{
			banks: [][]int{
				{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
				{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
				{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
				{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			},
			expected: 357,
		},
		{
			banks: [][]int{
				{1, 1, 1, 9},
			},
			expected: 19,
		},
		{
			banks: [][]int{
				{1, 1, 8, 9},
			},
			expected: 89,
		},
		{
			banks: [][]int{
				{9, 8, 7, 1, 7, 8, 9},
			},
			expected: 99,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, puzzle3.Run(tc.banks))
		})
	}
}
