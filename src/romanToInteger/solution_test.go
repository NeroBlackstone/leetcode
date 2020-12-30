package romanToInteger

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "Case2",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "Case3",
			args: args{
				s: "IX",
			},
			want: 9,
		},
		{
			name: "Case4",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "Case5",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
		{
			name: "Case6",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "Case7",
			args: args{
				s: "WWWW",
			},
			want: 0,
		},
		{
			name: "Case8",
			args: args{
				s: "IIIIIIIIIIIIIIIIIIIIIIII",
			},
			want: 0,
		},
		{
			name: "Case9",
			args: args{
				s: "D",
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
