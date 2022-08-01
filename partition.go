package main

import (
	"fmt"
)

func main() {
	fmt.Println(Partitions(1))
}

func Partitions(n int) int {
	list := CreateSlice(n)
	if n == 0 {
		return 0
	}
	var res = 1
	for len(list) > 1 {
		minInd := Min(list)
		list[Min(list)] += 1
		list[len(list)-1] -= 1
		s := Sum(list[minInd+1:])
		list = append(list[:minInd+1], CreateSlice(s)...)
		res++
	}
	return res
}

func Min(list []int) (index int) {
	min := list[0]
	minIndex := 0
	for i, _ := range list[:len(list)-1] {
		if min > list[i] {
			min = list[i]
			minIndex = i
		}
	}
	return minIndex
}

func Sum(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func CreateSlice(n int) []int {
	list := make([]int, n)
	for i, _ := range list {
		list[i] = 1
	}
	return list
}
