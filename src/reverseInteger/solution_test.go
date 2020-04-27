package reverseInteger

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Case 1",
			args{123},
			321,
		},
		{
			"Case 2",
			args{-123},
			-321,
		},
		{
			"Case 3",
			args{120},
			21,
		},
		{
			"Case 4",
			args{1534236469},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
