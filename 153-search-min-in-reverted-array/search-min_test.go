package searchmininrevertedarray

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "case 1",
		args: args{nums: []int{3, 4, 5, 1, 2}},
		want: 1,
	},
	{
		name: "case 2",
		args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}},
		want: 0,
	},
	{
		name: "case 3",
		args: args{nums: []int{11, 13, 15, 17}},
		want: 11,
	},
	{
		name: "case 4",
		args: args{nums: []int{}},
		want: math.MinInt64,
	},
	{
		name: "case 5",
		args: args{nums: nil},
		want: math.MinInt64,
	},
}

func Test_findMin(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMin(tt.args.nums))
		})
	}
}
