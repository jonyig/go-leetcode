package algorithm

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {

	list := []int{5, 10, 2, 7, 1}
	//list := []int{55, 44, 90, 88, 77, 99, 11, 66, 33, 22}
	//list := []int{69, 81, 30, 38, 9, 2, 47, 61, 32, 79}

	resultDesc := HeapSortDesc(list)

	expectationDesc := []int{10, 7, 5, 2, 1}

	if !reflect.DeepEqual(resultDesc, expectationDesc) {
		t.Fatalf("expect is %d in %d, but fail", resultDesc, expectationDesc)
	}

	result := HeapSort(list)

	expectation := []int{1, 2, 5, 7, 10}

	if !reflect.DeepEqual(result, expectation) {
		t.Fatalf("expect is %d in %d, but fail", result, expectation)
	}

}
