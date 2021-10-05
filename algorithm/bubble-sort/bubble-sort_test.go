package algorithm

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{99, 90, 88, 77, 66, 55, 44, 33, 22, 11}

	result := BubbleSort(list)

	expectation := []int{11, 22, 33, 44, 55, 66, 77, 88, 90, 99}

	if !reflect.DeepEqual(result, expectation) {
		t.Fatalf("expect is %d in %d, but fail", expectation, result)
	}
}
