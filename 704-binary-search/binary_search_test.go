package binarysearch

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
	want int
}{
	{
		name: "case 1",
		args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 3},
		want: 2,
	},
	{
		name: "case 2",
		args: args{nums: []int{1, 0, 3, 5, 9, 12}, target: 2},
		want: -1,
	},
	{
		name: "case 3",
		args: args{nums: nil, target: 1},
		want: -1,
	},
	{
		name: "case 4",
		args: args{nums: []int{}, target: 1},
		want: -1,
	},
}

func Test_search(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, search(tt.args.nums, tt.args.target))
		})
	}
}
