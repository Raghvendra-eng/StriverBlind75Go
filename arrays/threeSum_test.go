package arrays

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]int
		want1 [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want:  [][]int{{-1, -1, 2}, {-1, 0, 1}},
			want1: [][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			name: "test2",
			args: args{
				nums: []int{-2, 0, 1, 1, 2},
			},
			want:  [][]int{{-2, 0, 2}, {-2, 1, 1}},
			want1: [][]int{{-2, 1, 1}, {-2, 0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
