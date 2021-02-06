package simplifyPath

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Case1",
			args{"/home/"},
			"/home",
		},
		{
			"Case2",
			args{"/../"},
			"/",
		},
		{
			"Case3",
			args{"/home//foo/"},
			"/home/foo",
		},
		{
			"Case",
			args{"/a/./b/../../c/"},
			"/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
