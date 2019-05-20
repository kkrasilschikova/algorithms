package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	var arr1 = []int{6, 5, 4, 3}
	var arr2 = []int{9, 2, 8, 6}
	var arr3 = []int{1, 2, 3}
	var arr4 []int

	var res1 = SelectionSort(arr1)
	if !reflect.DeepEqual(res1, []int{3, 4, 5, 6}) {
		t.Errorf("Result of sorting array %v is incorrect: %v, should be {3, 4, 5, 6}", arr1, res1)
	}
	var res2 = SelectionSort(arr2)
	if !reflect.DeepEqual(res2, []int{2, 6, 8, 9}) {
		t.Errorf("Result of sorting array %v is incorrect: %v, should be {2, 6, 8, 9}", arr2, res2)
	}
	var res3 = SelectionSort(arr3)
	if !reflect.DeepEqual(res3, []int{1, 2, 3}) {
		t.Errorf("Result of sorting array %v is incorrect: %v, should be {1, 2, 3}", arr3, res3)
	}
	var res4 = SelectionSort(arr4)
	if !reflect.DeepEqual(res4, arr4) {
		t.Errorf("Result of sorting array %v is incorrect: %v, should be []", arr4, res4)
	}
}
