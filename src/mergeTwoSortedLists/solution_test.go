package mergeTwoSortedLists

import (
	"github.com/NeroBlackstone/leetcode/collection"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *collection.ListNode
		l2 *collection.ListNode
	}

	tests := []struct {
		name string
		args args
		want *collection.ListNode
	}{
		{
			name: "Case1",
			args: args{
				l1: collection.NewLinkList([]int{1, 2, 4}),
				l2: collection.NewLinkList([]int{1, 3, 4}),
			},
			want: collection.NewLinkList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "Case2",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			name: "Case3",
			args: args{
				l1: nil,
				l2: collection.NewLinkList([]int{0}),
			},
			want: collection.NewLinkList([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
