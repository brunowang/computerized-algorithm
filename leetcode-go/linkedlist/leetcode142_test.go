package linkedlist

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	list := NewList(1, 2, 3, 4, 5)
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test", args{list.Copy()}, nil},
		{"test", args{list.LinkSelf(4, 2)}, list.Find(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
