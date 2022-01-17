package fibonaccinumber

import "testing"

func Test_fibBottomUp(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{1},
			want: 1,
		},
		{
			name: "2",
			args: args{2},
			want: 1,
		},
		{
			name: "3",
			args: args{3},
			want: 2,
		},
		{
			name: "4",
			args: args{9},
			want: 34,
		},
		{
			name: "5",
			args: args{4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibBottomUp(tt.args.n); got != tt.want {
				t.Errorf("fibBottomUp() = %v, want %v", got, tt.want)
			}
			if got := fibMemo(tt.args.n); got != tt.want {
				t.Errorf("fibMemo() = %v, want %v", got, tt.want)
			}
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
