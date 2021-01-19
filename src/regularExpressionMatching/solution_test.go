package regularExpressionMatching

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
