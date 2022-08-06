package main

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"sort"
	"strings"
)

const (
	columnB    = 1
	columnD    = 3
	st3        = "ст3"
	oc         = "оц"
	steel09g2s = "09г2с"
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
	workingRange := make(map[string]string, 0)
	var tmp string
	var outSlice []string
	rangeLow := 0
	rangeHigh := 0
	fmt.Println("Enter range of rows to work with")
	fmt.Scan(&rangeLow, &rangeHigh)
	var steelParam, thicknessParam string
	fmt.Println("Enter what kind of steel do you want to work with: (ст3, оц, 09г2с)")
	fmt.Scan(&steelParam)
	fmt.Println("Enter sheet thickness: (1,1.5,2,2.5,3,4)")
	fmt.Scan(&thicknessParam)
	if steelParam == st3 {
		for i := rangeLow; i < rangeHigh; i++ {
			if strings.Contains(sh.Cell(i, columnB).String(), st3) {
				workingRange[sh.Cell(i, columnB).String()] = sh.Cell(i, columnD).String()
			}
		}
	} else if steelParam == oc {
		for i := rangeLow; i < rangeHigh; i++ {
			if strings.Contains(sh.Cell(i, columnB).String(), oc) {
				workingRange[sh.Cell(i, columnB).String()] = sh.Cell(i, columnD).String()
			}
		}
	} else {
		err := errors.New("wrong sheet material input")
		fmt.Println(err)
	}

	for k, v := range workingRange {
		if strings.Contains(k, st3) && strings.Contains(v, thicknessParam) {
			space := strings.Index(k, " ")
			parenthesis := strings.Index(k, "(")
			tmp = k[:space] + " " + thicknessParam + "мм чернуха" + " на " + k[parenthesis+1:]
			tmp = strings.Trim(tmp, "№")
			tmp = strings.Trim(tmp, "ООО ЭС")
			tmp = strings.Trim(tmp, ")") + "\n"
			outSlice = append(outSlice, tmp)
		} else if strings.Contains(k, oc) && strings.Contains(v, thicknessParam) {
			space := strings.Index(k, " ")
			parenthesis := strings.Index(k, "(")
			tmp = k[:space] + " " + thicknessParam + "мм ОЦ" + " на " + k[parenthesis+1:]
			tmp = strings.Trim(tmp, "№")
			tmp = strings.Trim(tmp, "ООО ЭС")
			tmp = strings.Trim(tmp, ")") + "\n"
			outSlice = append(outSlice, tmp)
		}
	}
	sort.Strings(outSlice)
	fmt.Printf("%v", outSlice)
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, v := range outSlice {
		file.WriteString(v)
	}
}
