package bintree

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{&TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
		}}, []int{1, 2, 3}},
		{"test", args{nil}, nil},
		{"test", args{&TreeNode{Val: 1}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
