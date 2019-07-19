package dynamic

import (
	"reflect"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	inputWord := "fosh"
	dict := []string{"fish", "vista"}

	res := LongestCommonSubsequence(inputWord, dict)
	if !reflect.DeepEqual(res, "fish") {
		t.Errorf("Incorrect word, should be %v", res)
	}
}
