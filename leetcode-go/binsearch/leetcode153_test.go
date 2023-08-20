package binsearch

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{}}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}}, 1},
		{"test", args{[]int{1, 3, 3, 3, 5}}, 1},
		{"test", args{[]int{5, 7, 9, 1, 3}}, 1},
		{"test", args{[]int{5, 7, 9, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
