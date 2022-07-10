package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(DigitalRoot(0))
}

func DigitalRoot(n int) int {
	nStr := strconv.Itoa(n)
	nList := strings.Split(nStr, "")
	var resultSum int
	for _, v := range nList {
		a, _ := strconv.Atoi(v)
		resultSum += a
	}
	if resultSum > 10 {
		return 0 + DigitalRoot(resultSum)
	} else {
		return resultSum
	}
}
