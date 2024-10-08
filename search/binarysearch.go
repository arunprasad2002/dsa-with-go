package search

func BinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
