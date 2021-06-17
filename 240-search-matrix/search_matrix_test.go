package searchmatrixii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	matrix [][]int
	target int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "case 1",
		args: args{matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, target: 5},
		want: true,
	},
	{
		name: "case 2",
		args: args{matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, target: 20},
		want: false,
	},
	{
		name: "case 3",
		args: args{matrix: [][]int{{1}}, target: 1},
		want: true,
	},
	{
		name: "case 4",
		args: args{matrix: [][]int{{1, 2}}, target: 3},
		want: false,
	},
	{
		name: "case 5",
		args: args{matrix: [][]int{{1, 2}}, target: 2},
		want: true,
	},
	{
		name: "case 6",
		args: args{matrix: [][]int{{1}, {2}}, target: 3},
		want: false,
	},
	{
		name: "case 7",
		args: args{matrix: [][]int{{1}, {2}}, target: 1},
		want: true,
	},
	{
		name: "case 8",
		args: args{matrix: [][]int{}, target: 1},
		want: false,
	},
	{
		name: "case 9",
		args: args{matrix: nil, target: 1},
		want: false,
	},
}

func Test_searchMatrix(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchMatrix(tt.args.matrix, tt.args.target))
		})
	}
}

func Test_searchMatrixV2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchMatrixV2(tt.args.matrix, tt.args.target))
		})
	}
}

func Benchmark_searchMatrixV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchMatrixV2(tests[0].args.matrix, tests[0].args.target)
	}
}

func Benchmark_searchMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchMatrix(tests[0].args.matrix, tests[0].args.target)
	}
}
