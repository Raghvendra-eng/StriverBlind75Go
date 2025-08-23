package binary

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args {
				a: 0,
				b: 0,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args {
				a: 5,
				b: 4,
			},
			want: 9,
		},
		{
			name: "test3",
			args: args {
				a: -5,
				b: 4,
			},
			want: -1,
		},
		{
			name: "test4",
			args: args {
				a: -1,
				b: 1,
			},
			want: 0,
		},
		{
			name: "test5",
			args: args {
				a: 3,
				b: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
