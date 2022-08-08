package main

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	columnB    = 1
	columnD    = 3
	st3        = "ст3"
	oc         = "оц"
	steel09g2s = "09Г2С"
	exit       = "выход"
)

func main() {
	// open an existing file
	wb, err := xlsx.OpenFile("Приоритетность архивов на координатный станок.xlsx")
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
	fmt.Println("Введите диапазон для обработки:")
	_, err = fmt.Scan(&rangeLow, &rangeHigh)
	if err != nil {
		fmt.Println("Проверьте типы входных параметров")
		return
	}
	if rangeLow > rangeHigh {
		fmt.Println(errors.New("ошибка: нижняя граница не может быть больше верхней"))
		return
	}
	var steelParam, thicknessParam string
	steelParam, err = inputSteel()
	if steelParam == exit {
		return
	}
	fmt.Println("Введите толщину листа: (1,2,3,4)")
	fmt.Scan(&thicknessParam)
	converted, err := strconv.Atoi(thicknessParam)
	if err != nil {
		return
	}
	if converted > 4 || converted < 1 {
		fmt.Println(errors.New("ошибка: заданная толщина не обрабатывается, либо введена неверно"))
		return
	}
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
	} else if steelParam == steel09g2s {
		for i := rangeLow; i < rangeHigh; i++ {
			if strings.Contains(sh.Cell(i, columnB).String(), steel09g2s) {
				workingRange[sh.Cell(i, columnB).String()] = sh.Cell(i, columnD).String()
			}
		}
	} else {
		err := errors.New("неправильно введён материал")
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
		} else if strings.Contains(k, steel09g2s) && strings.Contains(v, thicknessParam) {
			space := strings.Index(k, " ")
			parenthesis := strings.Index(k, "(")
			tmp = k[:space] + " " + thicknessParam + "мм 09г2с" + " на " + k[parenthesis+1:]
			tmp = strings.Trim(tmp, "№")
			tmp = strings.Trim(tmp, "ООО ЭС")
			tmp = strings.Trim(tmp, ")") + "\n"
			outSlice = append(outSlice, tmp)
		}
	}
	sort.Strings(outSlice)
	if len(outSlice) == 0 {
		errEmptyOut := errors.New("в данном диапазоне металла задонного типа не обнаружено")
		fmt.Println(errEmptyOut)
	}
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, v := range outSlice {
		file.WriteString(v)
	}
}

//Функция inputSteel сканирует ввод стали и если он не соответствуем трём вариантам, выводит ошибку и вызывает
//функцию по новой пока не будетт введён верный
func inputSteel() (string, error) {
	var steelParam string
	fmt.Println("Введите сталь, с которой хотите работать: (ст3, оц, 09г2с)")
	fmt.Println("Введите \"выход\" для выхода из программы")
	fmt.Scan(&steelParam)
	switch steelParam {
	case st3:
		return st3, nil
	case oc:
		return oc, nil
	case steel09g2s:
		return steel09g2s, nil
	case exit:
		return exit, nil
	default:
		fmt.Println(errors.New("ошибка: неверно введён материал"))
		return inputSteel()
	}
}
