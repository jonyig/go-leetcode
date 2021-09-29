package algorithm

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {

	//list := []int{1,17,22,31,43,56,63,76,88,90,100}
	var list []int

	for i := 1; i <= 10000; i++ {
		list = append(list, i)
	}
	target := 789
	result := BinarySearch(list, target)
	expect := 788

	if result != expect {
		t.Fatalf("expect is %d in %d, but faitl", target, expect)
	}
}
