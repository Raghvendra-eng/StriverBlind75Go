package dynamicprogamming

import "testing"

func Test_climbStairs(t *testing.T) {
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
			name: "test1",
			args: args{2},
			want: 2,
		},
		{
			name: "test1",
			args: args{3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
			if got := climbStairs1(tt.args.n); got != tt.want {
				t.Errorf("climbStairs1() = %v, want %v", got, tt.want)
			}
			if got := climbStairs2(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
