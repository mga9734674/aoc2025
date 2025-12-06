package puzzle4_test

import (
	"strconv"
	"testing"

	"aoc2025/internal/puzzle4"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	type testCase struct {
		rows     [][]bool
		expected int
	}

	testCases := []testCase{
		{
			rows: [][]bool{
				{false, false, true, true, false, true, true, true, true, false},
				{true, true, true, false, true, false, true, false, true, true},
				{true, true, true, true, true, false, true, false, true, true},
				{true, false, true, true, true, true, false, false, true, false},
				{true, true, false, true, true, true, true, false, true, true},
				{false, true, true, true, true, true, true, true, false, true},
				{false, true, false, true, false, true, false, true, true, true},
				{true, false, true, true, true, false, true, true, true, true},
				{false, true, true, true, true, true, true, true, true, false},
				{true, false, true, false, true, true, true, false, true, false},
			},
			expected: 13,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, puzzle4.Run(tc.rows))
		})
	}
}
