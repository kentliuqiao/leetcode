package searchrange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	nums   []int
	target int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "case 1",
		args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 8},
		want: []int{3, 4},
	},
	{
		name: "case 2",
		args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 6},
		want: []int{-1, -1},
	},
	{
		name: "case 3",
		args: args{nums: []int{}, target: 1},
		want: []int{-1, -1},
	},
	{
		name: "csae 4",
		args: args{nums: nil, target: 1},
		want: []int{-1, -1},
	},
	{
		name: "case 5",
		args: args{nums: []int{-1, 2, 3, 4, 4, 6}, target: 6},
		want: []int{5, 5},
	},
}

func Test_searchRange(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchRange(tt.args.nums, tt.args.target))
		})
	}
}
