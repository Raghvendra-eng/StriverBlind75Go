package arrays

import "testing"

func Test_maximumProfit(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr: []int{7, 1, 5, 3, 6},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProfit(tt.args.arr); got != tt.want {
				t.Errorf("maximumProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
