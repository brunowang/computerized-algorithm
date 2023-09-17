package bintree

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
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
				1,
				&TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 4}},
				&TreeNode{5, nil, &TreeNode{Val: 6}},
			},
		}, &TreeNode{
			Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{
				Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
			}}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatten(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
