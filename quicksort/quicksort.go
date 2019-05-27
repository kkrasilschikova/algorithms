package quicksort

/*
QuickSort returns a sorted array
*/
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	{
		pivot := arr[0]
		var less []int
		var greater []int

		for _, v := range arr[1:] {
			if v <= pivot {
				less = append(less, v)
			}
		}
		for _, v := range arr[1:] {
			if v > pivot {
				greater = append(greater, v)
			}
		}
		return append(append(QuickSort(less), pivot), QuickSort(greater)...)
	}
}
