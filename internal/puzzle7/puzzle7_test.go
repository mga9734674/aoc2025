package puzzle7_test

import (
	"bytes"
	"strconv"
	"testing"

	"aoc2025/internal/puzzle7"

	"github.com/stretchr/testify/require"
)

const test3 = `
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

func TestRun(t *testing.T) {
	type testCase struct {
		x        []byte
		expected int
	}

	testCases := []testCase{
		{
			x:        bytes.TrimPrefix([]byte(test3), []byte("\n")),
			expected: 40,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := puzzle7.Run(tc.x)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}
