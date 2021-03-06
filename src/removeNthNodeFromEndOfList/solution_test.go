package removeNthNodeFromEndOfList

import (
	"reflect"
	"testing"

	"github.com/NeroBlackstone/leetcode/collection"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *collection.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *collection.ListNode
	}{
		{
			name: "Case1",
			args: args{
				head: collection.NewLinkedList([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: collection.NewLinkedList([]int{1, 2, 3, 5}),
		},
		{
			name: "Case2",
			args: args{
				head: collection.NewLinkedList([]int{1}),
				n:    1,
			},
			want: collection.NewLinkedList([]int{}),
		},
		{
			name: "Case3",
			args: args{
				head: collection.NewLinkedList([]int{1, 2}),
				n:    1,
			},
			want: collection.NewLinkedList([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
