package stringToInteger

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{"42"},
			42,
		},
		{
			"case 2",
			args{"   -42"},
			-42,
		},
		{
			"case 3",
			args{"4193 with words"},
			4193,
		},
		{
			"case 4",
			args{"-91283472332"},
			-2147483648,
		},
		{
			"case 5",
			args{"words and 987"},
			0,
		},
		{
			"case 6",
			args{"3.14159"},
			3,
		},
		{
			"case 7",
			args{"  -0012a42"},
			-12,
		},
		{
			"case 8",
			args{"+1"},
			1,
		},
		{
			"case 9",
			args{"2147483648"},
			2147483647,
		},
		{
			"case 10",
			args{"20000000000000000000"},
			2147483647,
		},
		{
			"case 11",
			args{"-5-"},
			-5,
		},
		{
			"case 12",
			args{"-13+8"},
			-13,
		},
		{
			"case 13",
			args{""},
			0,
		},
		{
			"case 14",
			args{" "},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invalidRune(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case 1",
			args{' '},
			true,
		},
		{
			"case 2",
			args{'+'},
			false,
		},
		{
			"case 3",
			args{'e'},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invalidRune(tt.args.r); got != tt.want {
				t.Errorf("invalidRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isANumber(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case 1",
			args{' '},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isANumber(tt.args.r); got != tt.want {
				t.Errorf("isANumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
