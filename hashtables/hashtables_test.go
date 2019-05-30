package hashtables

import (
	"reflect"
	"testing"
)

func TestHashTables(t *testing.T) {
	var arr = []int{1, 15, 26, 33, 48, 100}
	var hashRes = map[int][]int{
		1: {1, 15},
		2: {100},
		5: {26, 33},
		6: {48},
	}

	res := HashTable(arr, 7)
	if !reflect.DeepEqual(res, hashRes) {
		t.Errorf("Hash table is incorrect, resulting hash table should be %v", res)
	}
}
