package longestCommonPrefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case1",
			args: args{
				[]string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "Case2",
			args: args{
				[]string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "Case3",
			args: args{
				[]string{},
			},
			want: "",
		},
		{
			name: "Case3",
			args: args{
				[]string{""},
			},
			want: "",
		},
		{
			name: "Case4",
			args: args{
				[]string{"a"},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
