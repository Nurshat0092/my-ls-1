package main

import (
	"fmt"
	"os"
)

func main() {
	printWrittenFiles()
	m.Files = []string{}
	if len(m.Dirs) > 0 {
		solve(m.Dirs[0])
		for _, w := range m.Dirs[1:] {
			solve(w)
		}
	}
}

func solve(pathStr string) {
	f, err := os.Open(pathStr)
	if os.IsPermission(err) {
		fmt.Println("No permission to enter file pathStr")
		return
	}
	subFiles, err := f.Readdir(-1)
	if os.IsPermission(err) {
		fmt.Println("No permission to enter file pathStr")
		return
	}
	if !m.Fa {
		subFiles = removeHidden(subFiles)
	} else {
		d, _ := os.Lstat(pathStr + "/.")
		dd, _ := os.Lstat(pathStr + "/..")
		subFiles = append(subFiles, []os.FileInfo{d, dd}...)
	}
	sortFiles(subFiles)
	if pathStr != m.Dirs[0] || len(m.Files) > 0 {
		fmt.Println()
	}
	if m.FR || len(m.Dirs) > 1 {
		fmt.Println(pathStr + ":")
	}
	if m.Fl {
		sizeLen, linkLen, totalM, giLen, uiLen := getAllMax(subFiles)
		fmt.Printf("total %d\n", totalM)
		printLong(subFiles, &pathStr, sizeLen, linkLen, totalM, giLen, uiLen)
	} else {
		printUsual(subFiles)
	}
	if m.FR {
		for _, w := range subFiles {
			if w.Mode().String()[0] == 'd' && w.Name() != "." && w.Name() != ".." {
				solve(pathStr + "/" + w.Name())
			}
		}
	}
}

func handleError(err *error) {
	if *err != nil {
		fmt.Println((*err).Error())
	}
}
