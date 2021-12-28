package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		}, {
			name: "   fly me   to   the moon  ",
			args: args{"   fly me   to   the moon  "},
			want: 4,
		}, {
			name: "luffy is still joyboy",
			args: args{"luffy is still joyboy"},
			want: 6,
		}, {
			name: "a",
			args: args{"a"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LengthOfLastWord(tt.args.s))
		})
	}
}
