package dynamic

/*
LongestCommonSubsequence returns a word that has the longest common subseauence with input
*/
func LongestCommonSubsequence(inputWord string, dict []string) string {
	maxLength := 0
	similarWord := ""

	for _, wordFromDict := range dict {
		grid := make([][]int, len(inputWord))
		for k := range inputWord {
			grid[k] = make([]int, len(wordFromDict))
		}

		for i := range inputWord {
			for j := range wordFromDict {
				if inputWord[i] == wordFromDict[j] {
					if i == 0 || j == 0 {
						grid[i][j] = 1
					} else {
						grid[i][j] = grid[i-1][j-1] + 1
					}
				} else {
					if i == 0 || j == 0 {
						grid[i][j] = 0
					} else {
						if grid[i-1][j] > grid[i][j-1] {
							grid[i][j] = grid[i-1][j]
						} else {
							grid[i][j] = grid[i][j-1]
						}
					}
				}

				if grid[i][j] > maxLength {
					maxLength = grid[i][j]
					similarWord = wordFromDict
				}
			}
		}
	}
	return similarWord
}
