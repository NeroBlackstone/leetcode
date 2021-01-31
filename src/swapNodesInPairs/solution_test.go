package swapNodesInPairs

import (
	"reflect"
	"testing"

	"github.com/NeroBlackstone/leetcode/collection"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *collection.ListNode
	}
	tests := []struct {
		name string
		args args
		want *collection.ListNode
	}{
		{
			"Case1",
			args{collection.NewLinkList([]int{1,2,3,4}),},
			collection.NewLinkList([]int{2,1,4,3}),
		},
		{
			"Case2",
			args{collection.NewLinkList([]int{}),},
			collection.NewLinkList([]int{}),
		},
		{
			"Case3",
			args{collection.NewLinkList([]int{1}),},
			collection.NewLinkList([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
