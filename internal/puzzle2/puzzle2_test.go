package puzzle2_test

import (
	"strconv"
	"testing"

	"aoc2025/internal/puzzle2"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	type testCase struct {
		ranges   []puzzle2.Range
		expected int
	}

	testCases := []testCase{
		{
			ranges: []puzzle2.Range{
				{Start: 11, End: 22},
				{Start: 95, End: 115},
				{Start: 998, End: 1012},
				{Start: 1188511880, End: 1188511890},
				{Start: 222220, End: 222224},
				{Start: 1698522, End: 1698528},
				{Start: 446443, End: 446449},
				{Start: 38593856, End: 38593862},
			},
			expected: 1227775554,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, puzzle2.Run(tc.ranges))
		})
	}
}
