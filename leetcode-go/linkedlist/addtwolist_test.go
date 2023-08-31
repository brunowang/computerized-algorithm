package linkedlist

import (
	"reflect"
	"testing"
)

func Test_addTwoList(t *testing.T) {
	type args struct {
		listA *ListNode
		listB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test", args{NewList(6, 5, 8), NewList(3, 4, 7)}, NewList(1, 0, 0, 5)},
		{"test", args{NewList(6, 5, 8), NewList(4, 3, 4, 7)}, NewList(5, 0, 0, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoList(tt.args.listA, tt.args.listB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoList() = %v, want %v", got, tt.want)
			}
		})
	}
}
