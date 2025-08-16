package arrays

import (
	"reflect"
	"testing"
)

func Test_checkTwoSum(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr:    []int{2, 7, 11, 15},
				target: 9,
			},
			want:  true,
			want1: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkTwoSum(tt.args.arr, tt.args.target)
			if got != tt.want {
				t.Errorf("checkTwoSum() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("checkTwoSum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_checkTwoSumSortedArray(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr:    []int{2, 7, 11, 15},
				target: 9,
			},
			want:  true,
			want1: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkTwoSumSortedArray(tt.args.arr, tt.args.target)
			if got != tt.want {
				t.Errorf("checkTwoSumSortedArray() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("checkTwoSumSortedArray() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
