package leetcode59

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	n    int
	want [][]int
}{
	{
		n:    1,
		want: [][]int{{1}},
	},
	{
		n:    2,
		want: [][]int{{1, 2}, {4, 3}},
	},
	{
		n:    4,
		want: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
	},
}

func TestGenerateSpiralMatrix(t *testing.T) {
	for _, c := range cases {
		got := generateSpiralMatrix(c.n)
		assert.Equal(t, c.want, got)
	}
}
