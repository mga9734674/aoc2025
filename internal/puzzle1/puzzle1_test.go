package puzzle1_test

import (
	"strconv"
	"testing"

	"aoc2025/internal/puzzle1"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	type testCase struct {
		rotations []puzzle1.Rotation
		expected  int
	}

	testCases := []testCase{
		{
			rotations: []puzzle1.Rotation{
				{IsTrigo: true, N: 100},
			},
			expected: 1,
		},
		{
			rotations: []puzzle1.Rotation{
				{IsTrigo: false, N: 100},
			},
			expected: 1,
		},
		{
			rotations: []puzzle1.Rotation{
				{IsTrigo: true, N: 68},
				{IsTrigo: true, N: 30},
				{IsTrigo: false, N: 48},
				{IsTrigo: true, N: 5},
				{IsTrigo: false, N: 60},
				{IsTrigo: true, N: 55},
				{IsTrigo: true, N: 1},
				{IsTrigo: true, N: 99},
				{IsTrigo: false, N: 14},
				{IsTrigo: true, N: 82},
			},
			expected: 6,
		},
		{
			rotations: []puzzle1.Rotation{
				{IsTrigo: true, N: 168},   // 82 2
				{IsTrigo: true, N: 30},    // 52 2
				{IsTrigo: true, N: 152},   // 0 4
				{IsTrigo: false, N: 2},    // 2 4
				{IsTrigo: false, N: 298},  // 0 7
				{IsTrigo: false, N: 1},    // 1 7
				{IsTrigo: true, N: 101},   // 0 9
				{IsTrigo: true, N: 1000},  // 0 19
				{IsTrigo: false, N: 1000}, // 0 29
			},
			expected: 29,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, puzzle1.Run(tc.rotations))
		})
	}
}
