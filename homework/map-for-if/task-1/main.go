package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var resault []int
	var temp []int
	sort.Ints(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			resault = append(arr[i:])
		}
	}
	fmt.Println(resault, temp)
}
