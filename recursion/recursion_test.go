package recursion

import (
	"testing"
)

func TestListSum(t *testing.T) {
	list1 := []int{}
	list2 := []int{5}
	list3 := []int{1, 2, 3, 4, 5}

	var s1 = ListSum(list1)
	if s1 != 0 {
		t.Errorf("Sum of %v is incorrect %v, should be 0", list1, s1)
	}
	var s2 = ListSum(list2)
	if s2 != 5 {
		t.Errorf("Sum of %v is incorrect %v, should be 5", list2, s2)
	}
	var s3 = ListSum(list3)
	if s3 != 15 {
		t.Errorf("Sum of %v is incorrect %v, should be 15", list3, s3)
	}
}

func TestReverseString(t *testing.T) {
	str1 := ""
	str2 := "a"
	str3 := "cat"
	str4 := "live not on evil"

	var s1 = ReverseString(str1)
	if s1 != "" {
		t.Errorf("Reverse of %s is incorrect %s, should be ''", str1, s1)
	}
	var s2 = ReverseString(str2)
	if s2 != "a" {
		t.Errorf("Reverse of %s is incorrect %s, should be 'a'", str2, s2)
	}
	var s3 = ReverseString(str3)
	if s3 != "tac" {
		t.Errorf("Reverse of %s is incorrect %s, should be 'tac'", str3, s3)
	}
	var s4 = ReverseString(str4)
	if s4 != "livenotonevil" {
		t.Errorf("Reverse of %s is incorrect %s, should be 'livenotonevil'", str4, s4)
	}

}
