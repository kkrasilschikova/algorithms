package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var arr1 = []int{}
	var arr2 = []int{1}
	var arr3 = []int{9, 8, 7, 1, 2, 3}

	res1 := QuickSort(arr1)
	if !reflect.DeepEqual(res1, arr1) {
		t.Errorf("Result of sorting array %v is incorrect: %v, should be []", arr1, res1)
	}
	res2 := QuickSort(arr2)
	if !reflect.DeepEqual(res2, arr2) {
		t.Errorf("Result of sorting array %v is incorrect: %v, should be {1}", arr2, res2)
	}
	res3 := QuickSort(arr3)
	if !reflect.DeepEqual(res3, []int{1, 2, 3, 7, 8, 9}) {
		t.Errorf("Result of sorting array %v is incorrect: %v, should be {1, 2, 3, 7, 8, 9}", arr3, res3)
	}
}
