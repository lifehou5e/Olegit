package main

import "fmt"

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var resault []int
	for i, _ := range arr {
		if Contains(resault, arr[i]) == false {
			resault = append(resault, arr[i])
		}
	}
	fmt.Println(resault)
}
