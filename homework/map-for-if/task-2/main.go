package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const MaxInt32 = 1<<31 - 1
	input := "1 f 99 3 -1 -2 44 4 102"
	s := strings.Fields(input)
	var min, max int32
	min = MaxInt32
	for i := 0; i < len(s); i++ {
		char, err := strconv.ParseInt(s[i], 10, 32)
		if err != nil {
			continue
		}
		if int32(char) > max {
			max = int32(char)
			// fmt.Println("max =", max)
		}
		if int32(char) < min {
			min = int32(char)
			// fmt.Println("min = ", min)
		}
	}
	var resault string
	m := int(max)
	n := int(min)
	resault = strconv.Itoa(m) + " " + strconv.Itoa(n)
	fmt.Println(resault)
}
