package leetcode

import "testing"

func TestProblem(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Hello World",
			args: args{"Hello World"},
			want: 5,
		},{
			name: "   fly me   to   the moon  ",
			args: args{"   fly me   to   the moon  "},
			want: 4,
		},{
			name: "luffy is still joyboy",
			args: args{"luffy is still joyboy"},
			want: 6,
		},{
			name: "a",
			args: args{"a"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
}
