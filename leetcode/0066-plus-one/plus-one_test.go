package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblem(t *testing.T) {
	type args struct {
		s []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name:"[1,2,3]",
			args: args{[]int{1,2,3}},
			want: []int{1,2,4},
		},{
			name:"[4,3,2,1]",
			args: args{[]int{4,3,2,1}},
			want: []int{4,3,2,2},
		},{
			name:"[9]",
			args: args{[]int{9}},
			want: []int{1,0},
		},{
			name:"[9,9]",
			args: args{[]int{9,9}},
			want: []int{1,0,0},
		},{
			name:"[8,9,9,9]",
			args: args{[]int{8,9,9,9}},
			want: []int{9,0,0,0},
		},{
			name:"[9,8,9]",
			args: args{[]int{9,8,9}},
			want: []int{9,9,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, plusOne(tt.args.s))
		})
	}
}
