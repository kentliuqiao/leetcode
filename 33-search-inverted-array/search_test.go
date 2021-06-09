package searchinvertedarray

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
		args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
		want: 4,
	},
	{
		name: "case 2",
		args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
		want: -1,
	},
	{
		name: "case 3",
		args: args{nums: []int{}, target: 1},
		want: -1,
	},
	{
		name: "case 4",
		args: args{nums: []int{1}, target: 2},
		want: -1,
	},
	{
		name: "case 5",
		args: args{nums: nil, target: 1},
		want: -1,
	},
	{
		name: "case 7",
		args: args{nums: []int{4, 5, 0, 1, 2, 3}, target: 3},
		want: 5,
	},
	{
		name: "case 8",
		args: args{nums: []int{4, 5, 0, 1, 2, 3}, target: 6},
		want: -1,
	},
	{
		name: "case 9",
		args: args{nums: []int{7, 1, 2, 3, 4, 5}, target: 1},
		want: 1,
	},
	{
		name: "case 10",
		args: args{nums: []int{1, 2, 3, 4}, target: 3},
		want: 2,
	},
}

func Test_search(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, search(tt.args.nums, tt.args.target))
		})
	}
}
