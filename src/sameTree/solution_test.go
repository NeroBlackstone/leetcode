package sameTree

import (
	"github.com/NeroBlackstone/leetcode/collection"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *collection.TreeNode
		q *collection.TreeNode
	}

	// [1,2,3]
	c1a := &collection.TreeNode{
		Val: 1,
		Left: &collection.TreeNode{
			Val: 2,
		},
		Right: &collection.TreeNode{
			Val: 3,
		},
	}

	//[1,2]
	c2a1 := &collection.TreeNode{
		Val: 1,
		Left: &collection.TreeNode{
			Val: 2,
		},
	}

	//[1,null,2]
	c2a2 := &collection.TreeNode{
		Val: 1,
		Right: &collection.TreeNode{
			Val: 3,
		},
	}

	// [1,2,1]
	c3a1 := &collection.TreeNode{
		Val: 1,
		Left: &collection.TreeNode{
			Val: 2,
		},
		Right: &collection.TreeNode{
			Val: 1,
		},
	}

	// [1,1,2]
	c3a2 := &collection.TreeNode{
		Val: 1,
		Left: &collection.TreeNode{
			Val: 1,
		},
		Right: &collection.TreeNode{
			Val: 2,
		},
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case1",
			args: args{
				p: c1a,
				q: c1a,
			},
			want: true,
		},
		{
			name: "Case2",
			args: args{
				p: c2a1,
				q: c2a2,
			},
			want: false,
		},
		{
			name: "Case3",
			args: args{
				p: c3a1,
				q: c3a2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
