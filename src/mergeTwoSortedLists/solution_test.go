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

	//[1,2,4]
	c1l1 := &collection.ListNode{
		Val: 1,
		Next: &collection.ListNode{
			Val: 2,
			Next: &collection.ListNode{
				Val: 4,
			},
		},
	}
	// [1,3,4]
	c1l2 := &collection.ListNode{
		Val: 1,
		Next: &collection.ListNode{
			Val: 3,
			Next: &collection.ListNode{
				Val: 4,
			},
		},
	}
	//Output: [1,1,2,3,4,4]
	c1out := &collection.ListNode{
		Val: 1,
		Next: &collection.ListNode{
			Val: 1,
			Next: &collection.ListNode{
				Val: 2,
				Next: &collection.ListNode{
					Val: 3,
					Next: &collection.ListNode{
						Val: 4,
						Next: &collection.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}

	c3l2 := &collection.ListNode{
		Val: 0,
	}

	tests := []struct {
		name string
		args args
		want *collection.ListNode
	}{
		{
			name: "Case1",
			args: args{
				l1: c1l1,
				l2: c1l2,
			},
			want: c1out,
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
				l2: c3l2,
			},
			want: c3l2,
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
