package greedy

/*
Greedy returns a set of stations that cover the most states
*/
func Greedy(statesNeeded []string, stations map[string][]string) []string {
	var finalStations []string

	for len(statesNeeded) > 0 {
		var bestStation string
		var statesCovered []string

		for station, statesInStation := range stations {
			covered := makeUnique(intersect(statesNeeded, statesInStation))
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
				finalStations = append(finalStations, bestStation)
				statesNeeded = remove(statesNeeded, statesCovered)
			}
		}
	}

	return finalStations
}

func makeUnique(s []string) []string {
	var list []string
	for _, value := range s {
		if !contains(list, value) {
			list = append(list, value)
		}
	}
	return list
}

func contains(arr []string, elem string) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

func remove(s []string, toRemove []string) []string {
	var list []string
	for _, value := range s {
		if !contains(toRemove, value) {
			list = append(list, value)
		}
	}
	return list
}

func intersect(s1 []string, s2 []string) []string {
	var list []string
	for _, value := range s1 {
		if contains(s2, value) {
			list = append(list, value)
		}
	}
	return list
}
