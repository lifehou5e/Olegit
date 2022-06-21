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
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			resault = append(resault, arr[i])
		}
	}
	var lengthArr = len(arr)
	for i := 1; i < lengthArr-1; i++ {
		if arr[lengthArr-i] != arr[lengthArr-i-1] {
			temp = append(temp, arr[lengthArr-i])
		}
	}
	resault = append(resault, temp[0])
	fmt.Println(resault)
}
