package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

const (
	columnB = 1
	columnD = 3
	st3     = "ст3"
	oc      = "оц"
)

func main() {
	// open an existing file
	wb, err := xlsx.OpenFile("table.xlsx")
	if err != nil {
		panic(err)
	}
	sh, ok := wb.Sheet["Заказы в работе"]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}
	var workingRange []string
	var tmp string
	rangeLow := 0
	rangeHigh := 0
	fmt.Println("Enter range of rows to work with")
	fmt.Scan(&rangeLow, &rangeHigh)
	var steelParam, thicknessParam string
	fmt.Println("Enter what kind of steel do you want to work with: (ст3, оц, 09г2с)")
	fmt.Scan(&steelParam)
	fmt.Println("Enter what sheet thickness: (1,1.5,2,2.5,3,4)")
	fmt.Scan(&thicknessParam)

	//fmt.Println(rangeLow, rangeHigh, sh.Row(1))
	for i := rangeLow; i < rangeHigh; i++ {
		workingRange = append(workingRange, sh.Cell(i, columnB).String())
	}
	//fmt.Println(workingRange)
	for _, v := range workingRange {
		if strings.Contains(v, st3) {
			space := strings.Index(v, " ")
			parenthesis := strings.Index(v, "(")
			v = strings.Trim(v, ")")
			tmp = v[:space] + " " + thicknessParam + "мм" + " на " + v[parenthesis+1:]
			tmp = strings.Trim(tmp, "№")
		}
	}
	fmt.Println(tmp)
}
