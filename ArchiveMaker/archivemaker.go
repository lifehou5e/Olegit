package main

//C:\Users\lifeh\OneDrive\Documents\convert\olegit\ArchiveMaker
//№914(4423)ЭС

import (
	"fmt"
	"strings"
)

const (
	columnB          = 1
	columnC          = 2
	columnD          = 3
	startingRow      = 4
	ocinkovka        = "оц"
	steel3           = "сталь 3"
	pathToGeneralDir = "Z:\\ОБЩЕЕ\\ПРОЕКТЫ\\ЭС\\Заказы" //Z:\ОБЩЕЕ\ПРОЕКТЫ\ЭС\Заказы
)

type detail struct {
	detailName   string
	detailAmount string
}

type rangeRows struct {
	startRow int
	endRow   int
}

type MyStrSlice []string

func (m MyStrSlice) Index(s []string, str string) int {
	if strings.Contains(str, "3мм") && strings.Contains(str, ocinkovka) {
		for i, v := range s {
			if strings.Contains(v, "3 оц") {
				return i
			}
		}
	}
	if strings.Contains(str, "2,5мм") && strings.Contains(str, ocinkovka) {
		for i, v := range s {
			if strings.Contains(v, "2,5 оц") {
				return i
			}
		}
	}
	if strings.Contains(str, "2мм") && strings.Contains(str, ocinkovka) {
		for i, v := range s {
			if strings.Contains(v, "2 оц") {
				return i
			}
		}
	}
	if strings.Contains(str, "1,5мм") && strings.Contains(str, ocinkovka) {
		for i, v := range s {
			if strings.Contains(v, "1,5 оц") {
				return i
			}
		}
	}
	if strings.Contains(str, "1мм") && strings.Contains(str, ocinkovka) {
		for i, v := range s {
			if strings.Contains(v, "1 оц") {
				return i
			}
		}
	}
	if strings.Contains(str, "4мм") && strings.Contains(str, steel3) {
		for i, v := range s {
			if strings.Contains(v, "4 ст3") {
				return i
			}
		}
	}
	if strings.Contains(str, "3мм") && strings.Contains(str, steel3) {
		for i, v := range s {
			if strings.Contains(v, "3 ст3") {
				return i
			}
		}
	}
	if strings.Contains(str, "2мм") && strings.Contains(str, steel3) {
		for i, v := range s {
			if strings.Contains(v, "2 ст3") {
				return i
			}
		}
	}
	if strings.Contains(str, "1мм") && strings.Contains(str, steel3) {
		for i, v := range s {
			if strings.Contains(v, "1 ст3") {
				return i
			}
		}
	}

	return -1
}

func main() {

	pathToArchive, pathKB, pathToDirArchive, err := findDirArchive()
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := MakeBufferFromExcel(pathToArchive)
	if err != nil {
		return
	}

	filesPath, err := FindDrawings(pathKB, buffer)
	dirNameSlice, err := MakeDirs(buffer, pathToDirArchive)
	WriteDrawings(filesPath, dirNameSlice, pathToDirArchive, buffer)
	filesPath, err = FindDXF(pathKB, buffer)
	WriteDXF(filesPath, dirNameSlice, pathToDirArchive, buffer)

	fmt.Println("All good, i guess. You must check directories for mistakes.")
	fmt.Println("Enter any symbol and press Enter to exit.")

	var exit string
	fmt.Scan(&exit)
	if exit == "" {
		return
	}
}
