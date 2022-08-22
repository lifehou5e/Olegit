package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"strings"
)

func findDirArchive() (pathArchive string, pathToKB string, pathToDirArchive string, err error) {
	fmt.Printf("Trying to find archive in %s ...\n", pathToGeneralDir)
	pathTemp := make([]string, 0)
	var archive string
	fmt.Println("------------------")
	fmt.Println("Enter archive's number: (for example: 877-17)")
	fmt.Scan(&archive) //example: "877-1" - input

	order := strings.Split(archive, "-")[:1] //877

	pathTemp = append(pathTemp, pathToGeneralDir)

	files, err := os.ReadDir(pathToGeneralDir) //Z:\ОБЩЕЕ\ПРОЕКТЫ\ЭС\Заказы
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), order[0]) {
			pathTemp = append(pathTemp, file.Name()) //pathTemp = [pathToGeneralDir, "№877(4890) ООО ЭС"]
		}
	}
	d := pathTemp
	d = append(d, "КБ\\КД")
	pathToKB = strings.Join(d, "\\")
	pathTemp = append(pathTemp, "ТЕХ")                    // pathTemp = [pathToGeneralDir, "№877(4890) ООО ЭС", "ТЕХ"]
	pathTemp = append(pathTemp, "№"+order[0]+" "+"ВШТка") // pathTemp = [pathToGeneralDir, "№877(4890) ООО ЭС", "ТЕХ", "№877 ВШТка"]
	pathToArchive := strings.Join(pathTemp, "\\")

	files, err = os.ReadDir(pathToArchive) //C:\Users\lifeh\OneDrive\Documents\convert\olegit\ArchiveMaker\№877(4890) ООО ЭС\ТЕХ\№877 ВШТка

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		condSlice := strings.Split(file.Name(), " ")
		s := strings.Join(condSlice[:1], "")
		condSlice = strings.Split(s, "-")
		archiveSlice := strings.Split(archive, "-")

		if archiveSlice[1] == condSlice[1] {
			pathTemp = append(pathTemp, file.Name()) //[pathToGeneralDir, "№877(4890) ООО ЭС", "ТЕХ", "№877 ВШТка", "№877-1 ВШТ оц"]
			pathToDirArchive = strings.Join(pathTemp, "\\")
			pathTemp = append(pathTemp, file.Name()+".xlsx") //[pathToGeneralDir, "№877(4890) ООО ЭС", "ТЕХ", "№877 ВШТка", "№877-1 ВШТ оц", "№877-1 ВШТ оц.xlsx"]
		}
	}

	pathToArchive = strings.Join(pathTemp, "\\") //C:\Users\lifeh\OneDrive\Documents\convert\olegit\ArchiveMaker\№877(4890) ООО ЭС\ТЕХ\№877 ВШТка\№877-1 ВШТ оц\№877-1 ВШТ оц.xlsx
	fmt.Println("Archive successfully found.")
	return pathToArchive, pathToKB, pathToDirArchive, nil
}

func MakeBufferFromExcel(pathToArchive string) (buffer map[string][]detail) {
	fmt.Println("------------------")
	fmt.Println("Making buffer from Excel to work with.")
	wb, err := xlsx.OpenFile(pathToArchive)
	if err != nil {
		panic(err)
	}

	sh, ok := wb.Sheet["Лист1"]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}
	workTemp := make(map[string][]detail, 0)
	counterRange := make([]rangeRows, 0)
	i := startingRow
	countNotEmpty := 0
	for i < sh.MaxRow {
		if sh.Cell(i, columnD).String() == "" && sh.Cell(i, columnB).String() != "" {
			counterRange = append(counterRange, rangeRows{startRow: i})
			for sh.Cell(i, columnD).String() == "" && sh.Cell(i, columnB).String() != "" {

				i++
			}
			counterRange[countNotEmpty-1].endRow = i
			continue
		}
		i++
		countNotEmpty++
	}

	for _, v := range counterRange {
		i := v.startRow - 1
		workTemp[sh.Cell(i, columnD).String()] = append(workTemp[sh.Cell(i, columnD).String()], detail{sh.Cell(i, columnB).String(), sh.Cell(i, columnC).String()})
		c := 0 //iterator
		for j := v.startRow; c < v.endRow-v.startRow; c++ {
			workTemp[sh.Cell(i, columnD).String()] = append(workTemp[sh.Cell(i, columnD).String()], detail{sh.Cell(j, columnB).String(), sh.Cell(j, columnC).String()})
			j++
		}

	}

	fmt.Println("Done.")
	fmt.Println("------------------")

	return workTemp
}
