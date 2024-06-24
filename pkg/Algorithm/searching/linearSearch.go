// Copyright (c) avijit bhattacharjee 2024

package searching


import (
	"fmt"
)

func LinearSearch() {

	var input = []int{1,0,2,9,3,8,4,7,5,6}
	var search = 7
	fmt.Println("Element found at = ", linearSearch(input, search))

}

func linearSearch(arr []int, search int) int {

	for index, value := range arr {
		if value == search {
			return index
		}
	}
	return -1
}