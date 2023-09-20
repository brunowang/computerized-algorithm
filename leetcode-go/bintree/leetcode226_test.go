package bintree

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test", args{
			&TreeNode{
				4,
				&TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
				&TreeNode{7, &TreeNode{Val: 6}, &TreeNode{Val: 9}},
			},
		}, &TreeNode{
			4,
			&TreeNode{7, &TreeNode{Val: 9}, &TreeNode{Val: 6}},
			&TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 1}},
		}},
		{"test", args{
			&TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
		}, &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
