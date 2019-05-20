package selectionsort

/*
SelectionSort returns a sorted array
*/
func SelectionSort(arr []int) []int {
	var newArray []int
	length := len(arr)
	for i := 0; i < length; i++ {
		smallestIndex := findSmallest(arr)
		newArray = append(newArray, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return newArray
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}
