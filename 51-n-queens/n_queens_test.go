package nqueens

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "1", args: args{1}, want: [][]string{{"Q"}}},
		{name: "2", args: args{4}, want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
