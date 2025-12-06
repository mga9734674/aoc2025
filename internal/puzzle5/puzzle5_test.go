package puzzle5_test

import (
	"strconv"
	"testing"

	"aoc2025/internal/puzzle5"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	type testCase struct {
		ranges      []puzzle5.Range
		ingredients []int
		expected    int
	}

	testCases := []testCase{
		{
			ranges: []puzzle5.Range{
				{Start: 3, End: 5},
				{Start: 10, End: 14},
				{Start: 16, End: 20},
				{Start: 12, End: 18},
			},
			ingredients: []int{1, 5, 8, 11, 17, 32},
			expected:    3,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, puzzle5.Run(tc.ranges, tc.ingredients))
		})
	}
}
