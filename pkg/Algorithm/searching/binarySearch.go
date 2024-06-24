// Copyright (c) avijit bhattacharjee 2024

package searching

import (
	"fmt"
	"sort"
)

func BinarySearch() {
	
	var input = []int{0,1,2,3,4,5,6,7,8,9}
	var search = 7
	fmt.Println("Element found at = ", binarySearch(input, search, 0, len(input)-1))
}

func binarySearch(arr []int, search, start,end int) int {

	if start > end {
		return -1
	}

	var mid = (start+end)/2
	sort.Ints(arr)
	if arr[mid] == search {
		return mid
	}

	if search < arr[mid] {
		return binarySearch(arr, search, start, mid-1)
	}
	return binarySearch(arr, search, mid+1, end)
}