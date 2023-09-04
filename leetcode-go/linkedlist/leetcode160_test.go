package linkedlist

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	list := NewList(1, 2, 3, 4, 5, 6)
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test", args{NewList(1, 2, 3, 4, 5), list}, nil},
		{"test", args{NewList(1, 2).Link(list, 3), list}, list.Find(3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
