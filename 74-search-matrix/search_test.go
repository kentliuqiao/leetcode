package searchbidimensionalarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	nums   [][]int
	target int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "case 1",
		args: args{nums: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 3},
		want: true,
	},
	{
		name: "case 2",
		args: args{nums: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 13},
		want: false,
	},
	{
		name: "case 3",
		args: args{nums: [][]int{}, target: 1},
		want: false,
	},
	{
		name: "case 4",
		args: args{nums: [][]int{{}}, target: 1},
		want: false,
	},
	{
		name: "case 5",
		args: args{nums: [][]int{{1, 2, 3, 4}}, target: 3},
		want: true,
	},
}

func Test_searchMatrix(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchMatrix(tt.args.nums, tt.args.target))
		})
	}
}
