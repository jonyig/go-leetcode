package algorithm

import (
	"testing"
)

func TestSequentialSearch(t *testing.T) {

	var list []int
	for i := 1; i <= 10000; i++ {
		list = append(list, i)
	}
	target := 789
	result := SequentialSearch(list, target)

	expect := 788

	if result != expect {
		t.Fatalf("expect is %d in %d, but faitl", target, expect)
	}
}
