package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var list = []int{1, 2, 4, 7}

	var i1 = BinarySearch(list, 2)
	if i1 != 1 {
		t.Errorf("element 2 has index %d, should have index 1 in array %v", i1, list)
	}

	var i2 = BinarySearch(list, 9)
	if i2 != -1 {
		t.Errorf("element 9 has index %d, should have index -1 in array %v", i2, list)
	}
}
