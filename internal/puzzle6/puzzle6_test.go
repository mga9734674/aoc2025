package puzzle6_test

import (
	"strconv"
	"testing"

	"aoc2025/internal/puzzle6"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	type testCase struct {
		c        []puzzle6.Compute
		expected int
	}

	testCases := []testCase{
		{
			c: []puzzle6.Compute{
				{Nums: []int{123, 45, 6}, Ope: "*"},
				{Nums: []int{328, 64, 98}, Ope: "+"},
				{Nums: []int{51, 387, 215}, Ope: "*"},
				{Nums: []int{64, 23, 314}, Ope: "*+"},
			},
			expected: 4277556,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, puzzle6.Run(tc.c))
		})
	}
}
