package breadthfirstsearch

import (
	"reflect"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	var graph = map[string][]string{
		"A": {"B", "D"},
		"B": {"C", "E"},
		"C": {},
		"D": {"C", "F"},
		"E": {},
		"F": {"E"},
	}
	var start1 = "A"
	var finish1 = "F"
	var res1 = BreadthFirstSearch(graph, start1, finish1)
	if !reflect.DeepEqual(res1, true) {
		t.Error("Path is incorrect as path exists")
	}
	var start2 = "B"
	var finish2 = "F"
	var res2 = BreadthFirstSearch(graph, start2, finish2)
	if !reflect.DeepEqual(res2, false) {
		t.Error("Path is incorrect as path doesn't exist")
	}
}
