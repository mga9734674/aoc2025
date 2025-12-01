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
				{IsRight: false, N: 100},
			},
			expected: 0,
		},
		{
			rotations: []puzzle1.Rotation{
				{IsRight: true, N: 100},
			},
			expected: 0,
		},
		{
			rotations: []puzzle1.Rotation{
				{IsRight: false, N: 68},
				{IsRight: false, N: 30},
				{IsRight: true, N: 48},
				{IsRight: false, N: 5},
				{IsRight: true, N: 60},
				{IsRight: false, N: 55},
				{IsRight: false, N: 1},
				{IsRight: false, N: 99},
				{IsRight: true, N: 14},
				{IsRight: false, N: 82},
			},
			expected: 3,
		},
		{
			rotations: []puzzle1.Rotation{
				{IsRight: false, N: 168},
				{IsRight: false, N: 30},
				{IsRight: false, N: 152},
				{IsRight: true, N: 2},
				{IsRight: true, N: 298},
				{IsRight: true, N: 1},
				{IsRight: false, N: 101},
			},
			expected: 3,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, puzzle1.Run(tc.rotations))
		})
	}
}
