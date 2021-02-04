package linkedListCycle

import (
	"testing"

	"github.com/NeroBlackstone/leetcode/collection"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *collection.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Case1",
			args{collection.NewCircularLinkedList([]int{3,2,0,-4},1)},
			true,
		},
		{
			"Case2",
			args{collection.NewCircularLinkedList([]int{1,2},0)},
			true,
		},
		{
			"Case3",
			args{collection.NewCircularLinkedList([]int{1},-1)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
