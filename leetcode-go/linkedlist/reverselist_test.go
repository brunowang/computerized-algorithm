package linkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test", args{NewList([]int{})}, NewList([]int{})},
		{"test", args{NewList([]int{1})}, NewList([]int{1})},
		{"test", args{NewList([]int{1, 2})}, NewList([]int{2, 1})},
		{"test", args{NewList([]int{1, 2, 3})}, NewList([]int{3, 2, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
