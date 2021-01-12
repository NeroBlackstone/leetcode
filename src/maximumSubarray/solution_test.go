package maximumSubarray

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "Case2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Case3",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "Case4",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "Case5",
			args: args{
				nums: []int{-100000},
			},
			want: -100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
