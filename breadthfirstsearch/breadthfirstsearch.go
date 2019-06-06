package breadthfirstsearch

/*
BreadthFirstSearch finds the length os the shortest path between 2 points on graph
*/
func BreadthFirstSearch(graph map[string][]string, start string, finish string) bool {
	queue := graph[start]
	checked := []string{}
	for len(queue) > 0 {
		point := queue[0]
		if !contains(checked, point) {
			if point == finish {
				return true
			}
			{
				queue = append(queue[1:], graph[point]...)
				checked = append(checked, point)
			}
		} else {
			queue = queue[1:]
		}
	}
	return false
}

func contains(arr []string, elem string) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
