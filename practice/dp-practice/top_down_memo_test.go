package dppractice

import "testing"

// 1,1,2,3,5,8,13,21,34,55,89

func Test_fibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
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
			if got := fibonacci(tt.args.n); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
