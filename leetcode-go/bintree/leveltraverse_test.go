package bintree

import (
	"reflect"
	"testing"
)

func Test_levelTraverse(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 5}},
				Right: &TreeNode{3, &TreeNode{Val: 6}, &TreeNode{Val: 7}},
			},
		}, []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelTraverse(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
