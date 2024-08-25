package search

func LinearSearch(arr []int, target int) int {
	for i, number := range arr {
		if(target == number){
			return i
		}
	}

	return -1
}
