package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 2, 2, 4}
	var b = []int{2, 3, 5}
	fmt.Println(ArrayDiff(a, b))
}

func ArrayDiff(a, b []int) []int {
	result := make([]int, 0)
	m := map[int]bool{}
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if !m[v] {
			result = append(result, v)
		}
	}
	return result
}
