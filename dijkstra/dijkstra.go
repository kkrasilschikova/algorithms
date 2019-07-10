package dijkstra

/*
Dijkstra calculates the shortest path for a weighted graph
*/
func Dijkstra(graph map[string]map[string]int, costs map[string]int, parents map[string]string, processed []string) []string {
	node := FindLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbours := graph[node]

		nkeys := []string{}
		for key := range neighbours {
			nkeys = append(nkeys, key)
		}

		for _, n := range nkeys {
			newCost := cost + neighbours[n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed = append(processed, node)
		node = FindLowestCostNode(costs, processed)
	}
	return processed
}

/*
FindLowestCostNode is helper function that returns a node with the lowest cost
*/
func FindLowestCostNode(costs map[string]int, processed []string) string {
	lowestCost := 1000000
	var lowestCostNode string

	for node := range costs {
		cost := costs[node]
		if cost < lowestCost && !contains(processed, node) {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func contains(arr []string, elem string) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
