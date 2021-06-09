package searchinsert

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
		args: args{nums: []int{1, 3, 5, 6}, target: 5},
		want: 2,
	},
	{
		name: "case 2",
		args: args{nums: []int{1, 3, 5, 6}, target: 2},
		want: 1,
	},
	{
		name: "case 3",
		args: args{nums: []int{1, 3, 5, 6}, target: 7},
		want: 4,
	},
	{
		name: "case 4",
		args: args{nums: []int{1, 3, 5, 6}, target: 0},
		want: 0,
	},
	{
		name: "case 5",
		args: args{nums: []int{}, target: 1},
		want: 0,
	},
	{
		name: "case 6",
		args: args{nums: nil, target: 1},
		want: 0,
	},
}

func Test_searchInsert(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchInsert(tt.args.nums, tt.args.target))
		})
	}
}
