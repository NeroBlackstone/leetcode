package findFirstAndLastPositionOfElementInSortedArray

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Case1",
			args{[]int{5,7,7,8,8,10},8},
			[]int{3,4},
		},
		{
			"Case2",
			args{[]int{5,7,7,8,8,10},6},
			[]int{-1,-1},
		},
		{
			"Case3",
			args{[]int{},0},
			[]int{-1,-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
