package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindDXF(pathKB string, buffer map[string][]detail) (filesPath map[string][]string, err error) {
	fmt.Println("Finding .DXFs...")
	filesPath = make(map[string][]string)

	err = filepath.Walk(pathKB, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for key, val := range buffer {
			for i, _ := range val {
				if filepath.Ext(path) == ".DXF" && strings.Contains(os.FileInfo.Name(info), strings.TrimRight(buffer[key][i].detailName, " ")) {
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

func WriteDXF(filesPath map[string][]string, dirNameSlice MyStrSlice, pathToDirArchive string, buffer map[string][]detail) error {
	fmt.Println("Coping .DXFs into corresponding folders...")
	for metalThickness, paths := range filesPath {
		for index, _ := range paths {
			ind := dirNameSlice.Index(dirNameSlice, metalThickness)
			fin, err := os.Open(filesPath[metalThickness][index])
			if err != nil {
				log.Fatal(err)
			}
			defer fin.Close()

			fout, err := os.Create(pathToDirArchive + "\\раскрои" + "\\" + dirNameSlice[ind] + "\\" + buffer[metalThickness][index].detailName + " - " + buffer[metalThickness][index].detailAmount + " " + "шт." + ".DXF")
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
