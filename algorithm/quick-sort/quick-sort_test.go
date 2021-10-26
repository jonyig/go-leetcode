package algorithm

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []int{99, 90, 88, 77, 66, 55, 44, 33, 22, 11}
	//list := []int{55, 44, 90, 88, 77, 99, 11, 66, 33, 22}
	//list := []int{69, 81, 30, 38, 9, 2, 47, 61, 32, 79}

	result := QuickSort(list)

	expectation := []int{11, 22, 33, 44, 55, 66, 77, 88, 90, 99}

	if !reflect.DeepEqual(result, expectation) {
		t.Fatalf("expect is %d in %d, but fail", result, expectation)
	}
}
