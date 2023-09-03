package linkedlist

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	list00 := NewList(1, 2, 3, 4, 5)
	list01 := NewList(1, 2, 3, 4, 5, 6)
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test", args{list00}, list00.Find(2)},
		{"test", args{list01}, list01.Find(3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
