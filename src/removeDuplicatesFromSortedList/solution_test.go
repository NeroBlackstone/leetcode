package removeDuplicatesFromSortedList

import (
	"github.com/NeroBlackstone/leetcode/collection"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *collection.ListNode
	}
	tests := []struct {
		name string
		args args
		want *collection.ListNode
	}{
		{
			name: "Case1",
			args: args{
				head: collection.NewLinkedList([]int{1, 1, 2}),
			},
			want: collection.NewLinkedList([]int{1, 2}),
		},
		{
			name: "Case2",
			args: args{
				head: collection.NewLinkedList([]int{1, 1, 2, 3, 3}),
			},
			want: collection.NewLinkedList([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
