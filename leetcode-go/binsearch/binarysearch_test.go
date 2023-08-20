package binsearch

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{}, 0}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}, 0}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}, 10}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}, 1}, 0},
		{"test", args{[]int{1, 3, 3, 3, 5}, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leftBound(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{}, 0}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}, 0}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}, 10}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}, 1}, 0},
		{"test", args{[]int{1, 3, 3, 3, 5}, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftBound(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("leftBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightBound(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{}, 0}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}, 0}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}, 10}, -1},
		{"test", args{[]int{1, 3, 5, 7, 9}, 1}, 0},
		{"test", args{[]int{1, 3, 3, 3, 5}, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightBound(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("rightBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
