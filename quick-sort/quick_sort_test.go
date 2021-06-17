package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "case 1",
		args: args{nums: []int{2, 1, 3, -1, 20, 4}},
		want: []int{-1, 1, 2, 3, 4, 20},
	},
	{
		name: "case 2",
		args: args{nums: []int{1}},
		want: []int{1},
	},
	{
		name: "case 3",
		args: args{nums: []int{}},
		want: []int{},
	},
	{
		name: "case 4",
		args: args{nums: nil},
		want: nil,
	},
}

func Test_quickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, quickSort(tt.args.nums))
		})
	}
}

func Test_qsort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, qsort(tt.args.nums))
		})
	}
}
