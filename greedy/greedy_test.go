package greedy

import (
	"reflect"
	"testing"
)

func TestGreedy(t *testing.T) {
	var statesNeeded = []string{"mt", "wa", "or", "id", "nv", "ut", "az"}
	var stations = map[string][]string{
		"kone":   {"id", "nv", "ut"},
		"ktwo":   {"wa", "id", "mt"},
		"kthree": {"or", "nv", "ca"},
		"kfour":  {"nv", "ut"},
		"kfive":  {"ca", "az"},
	}

	var res = Greedy(statesNeeded, stations)
	if !reflect.DeepEqual(len(res), 4) {
		t.Errorf("Incorrect number of different stations %v", res)
	}
}
