package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestProblem(t *testing.T) {
    type args struct {
    		s int
    	}
    	tests := []struct {
    		name string
    		args args
    		want int
    	}{
			{
				name: "4",
				args: args{4},
				want: 2,
			},
			{
				name: "8",
				args: args{8},
				want: 2,
			},
			{
				name: "9",
				args: args{9},
				want: 3,
			},
    	}
    	for _, tt := range tests {
    		t.Run(tt.name, func(t *testing.T) {
    			assert.Equal(t, tt.want, mySqrt(tt.args.s))
    		})
    	}
}

func TestUseMathSqrt(t *testing.T) {
	for i:= 1;i<=100;i++ {
		_ = int(math.Sqrt(float64(i)))
	}
}
