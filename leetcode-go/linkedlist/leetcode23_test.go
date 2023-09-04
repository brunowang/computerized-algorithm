package linkedlist

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	lists := []*ListNode{NewList(1, 4, 5), NewList(1, 3, 4), NewList(2, 6)}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test", args{lists}, NewList(1, 1, 2, 3, 4, 4, 5, 6)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
