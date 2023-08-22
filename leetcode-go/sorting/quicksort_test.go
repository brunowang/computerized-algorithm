package sorting

import (
	"reflect"
	"testing"
)

func Test_quickSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{[]int{3, 5, 2, 7, 1}}, []int{1, 2, 3, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
