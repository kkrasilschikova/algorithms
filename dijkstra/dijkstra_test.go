package dijkstra

import (
	"reflect"
	"testing"
)

func TestFindLowestCostNode(t *testing.T) {
	var costs = map[string]int{
		"A":   6,
		"B":   2,
		"FIN": 100000,
	}
	var processed = []string{}

	var res = FindLowestCostNode(costs, processed)
	if !reflect.DeepEqual(res, "B") {
		t.Errorf("Lowest cost node is %v", res)
	}
}

func TestDijkstra(t *testing.T) {
	var graph = map[string]map[string]int{
		"START": {"A": 6, "B": 2},
		"A":     {"FIN": 1},
		"B":     {"A": 3, "FIN": 5},
		"FIN":   {},
	}
	var costs = map[string]int{
		"A":   6,
		"B":   2,
		"FIN": 100000,
	}
	var parents = map[string]string{
		"A":   "START",
		"B":   "START",
		"FIN": "",
	}
	var processed = []string{}

	var res = Dijkstra(graph, costs, parents, processed)
	if !reflect.DeepEqual(res, []string{"B", "A", "FIN"}) {
		t.Errorf("The shortest path from START to FIN is %v", res)
	}
}
