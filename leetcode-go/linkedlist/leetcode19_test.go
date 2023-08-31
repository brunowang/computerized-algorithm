package linkedlist

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test", args{NewList(1), 1}, NewList()},
		{"test", args{NewList(), 1}, NewList()},
		{"test", args{NewList(1, 2, 3), 3}, NewList(2, 3)},
		{"test", args{NewList(1, 2), 3}, NewList(1, 2)},
		{"test", args{NewList(1, 2, 3, 4), 1}, NewList(1, 2, 3)},
		{"test", args{NewList(1, 2, 3, 4), 2}, NewList(1, 2, 4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
