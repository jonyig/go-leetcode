package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblem(t *testing.T) {
    type args struct {
    		s int
    	}
    	tests := []struct {
    		name string
    		args int
    		want int
    	}{
    		{
				name:"1",
				args: 1,
				want: 1,
			},
			{
				name:"2",
				args: 2,
				want: 2,
			},{
				name:"3",
				args: 3,
				want: 3,
			},{
				name:"4",
				args: 4,
				want: 5,
			},{
				name:"5",
				args: 5,
				want: 8,
			},
    	}
    	for _, tt := range tests {
    		t.Run(tt.name, func(t *testing.T) {
    			assert.Equal(t, tt.want, climbStairs(tt.args))
    		})
    	}
}
