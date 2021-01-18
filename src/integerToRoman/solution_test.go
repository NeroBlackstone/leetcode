package integerToRoman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case1",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "Case2",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "Case3",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "Case4",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "Case5",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
