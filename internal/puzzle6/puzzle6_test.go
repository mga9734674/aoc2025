package puzzle6_test

import (
	"strconv"
	"testing"

	"aoc2025/internal/puzzle6"

	"github.com/stretchr/testify/require"
)

const test1 = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

const test2 = `945 486  2
778 257  4
488 182  9
785 344 72
+   *   + 
`

func TestRun(t *testing.T) {
	t.Parallel()

	type testCase struct {
		x        []byte
		expected int
	}

	testCases := []testCase{
		{
			x:        []byte(test1),
			expected: 3263827,
		},
		{
			x:        []byte(test2),
			expected: 243169394727,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			got, err := puzzle6.ParseAndRun(tc.x)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}
