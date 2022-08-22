package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func MakeDirs(buffer map[string][]detail, pathToDirArchive string) (dirNameSlice MyStrSlice, err error) {
	fmt.Println("Creating directories for sheet thicknesses...")
	if buffer == nil {
		return nil, errors.New("buffer is empty")
	}
	dirNameSlice = make(MyStrSlice, 0)

	for key, _ := range buffer {
		if strings.Contains(key, "1мм") && strings.Contains(key, ocinkovka) {
			dirNameSlice = append(dirNameSlice, "лист 1 оц")
		}
		if strings.Contains(key, "1,5мм") && strings.Contains(key, ocinkovka) {
			dirNameSlice = append(dirNameSlice, "лист 1,5 оц")
		}
		if strings.Contains(key, "2мм") && strings.Contains(key, ocinkovka) {
			dirNameSlice = append(dirNameSlice, "лист 2 оц")
		}
		if strings.Contains(key, "2,5мм") && strings.Contains(key, ocinkovka) {
			dirNameSlice = append(dirNameSlice, "лист 2,5 оц")
		}
		if strings.Contains(key, "3мм") && strings.Contains(key, ocinkovka) {
			dirNameSlice = append(dirNameSlice, "лист 3 оц")
		}
		if strings.Contains(key, "1мм") && strings.Contains(key, steel3) {
			dirNameSlice = append(dirNameSlice, "лист 1 ст3")
		}
		if strings.Contains(key, "2мм") && strings.Contains(key, steel3) {
			dirNameSlice = append(dirNameSlice, "лист 2 ст3")
		}
		if strings.Contains(key, "3мм") && strings.Contains(key, steel3) {
			dirNameSlice = append(dirNameSlice, "лист 3 ст3")
		}
		if strings.Contains(key, "4мм") && strings.Contains(key, steel3) {
			dirNameSlice = append(dirNameSlice, "лист 4 ст3")
		}
	}

	for _, v := range dirNameSlice {
		if err := os.Mkdir(pathToDirArchive+"\\раскрои\\"+v, 0755); err != nil {
			log.Fatal(err)
		}
		if err := os.Mkdir(pathToDirArchive+"\\чертежи\\"+v, 0755); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Directories for sheet thicknesses %v was created\n", dirNameSlice)
	fmt.Println("------------------")

	return dirNameSlice, nil
}

func FindDrawings(pathKB string, buffer map[string][]detail) (filesPath map[string][]string, err error) {
	fmt.Println("Searching for drawings...")
	filesPath = make(map[string][]string)

	err = filepath.Walk(pathKB, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for key, val := range buffer {
			for i, _ := range val {

				if filepath.Ext(path) == ".tif" && strings.Contains(os.FileInfo.Name(info), strings.TrimRight(buffer[key][i].detailName, " ")) {
					filesPath[key] = append(filesPath[key], path)
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("Done.")
	fmt.Println("------------------")

	return filesPath, nil
}

func WriteDrawings(filesPath map[string][]string, dirNameSlice MyStrSlice, pathToDirArchive string, buffer map[string][]detail) error {
	fmt.Println("Coping drawings into corresponding folders...")
	for metalThickness, paths := range filesPath {
		for index, _ := range paths {
			ind := dirNameSlice.Index(dirNameSlice, metalThickness)
			fin, err := os.Open(filesPath[metalThickness][index])
			if err != nil {
				log.Fatal(err)
			}
			defer fin.Close()

			fout, err := os.Create(pathToDirArchive + "\\чертежи" + "\\" + dirNameSlice[ind] + "\\" + buffer[metalThickness][index].detailName + ".tif")
			if err != nil {
				log.Fatal(err)
			}
			defer fout.Close()

			_, err = io.Copy(fout, fin)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("Done.")
	fmt.Println("------------------")

	return nil
}
