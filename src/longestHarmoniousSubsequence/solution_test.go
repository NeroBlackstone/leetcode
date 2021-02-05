package longestHarmoniousSubsequence

import "testing"

func Test_findLHS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Case1",
			args{[]int{1,3,2,2,5,2,3,7}},
			5,
		},
		{
			"Case2",
			args{[]int{1,2,3,4}},
			2,
		},
		{
			"Case3",
			args{[]int{1,1,1,1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLHS(tt.args.nums); got != tt.want {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
		})
	}
}
