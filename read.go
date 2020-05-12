package main

import (
	"fmt"
	"log"
	"os"
)

func (m *MyLS) readArgs(arr []string) {
	for i := 0; i < len(arr); i++ {
		if isFlag(arr[i]) {
			m.updateFlag(arr[i])
			arr[i] = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			i--
		}
	}
	for _, w := range arr {
		m.appendFile(w)
	}
	if len(arr) == 0 {
		m.Dirs = append(m.Dirs, ".")
	}
	sortFiles(m.Dirs)
	sortFiles(m.Files)
}

func (m *MyLS) updateFlag(str string) {
	str = str[1:]
	flags := []rune(str)
	for _, w := range flags {
		switch w {
		case 'a':
			m.Fa = true
		case 'R':
			m.FR = true
		case 'l':
			m.Fl = true
		case 'r':
			m.Fr = true
		case 't':
			m.Ft = true
		default:
			flagError := fmt.Sprintf("ls: invalid option -- '%s'\n", string(w))
			flagError += "Try 'ls --help' for more information."
			log.Fatalln(flagError)
		}
	}
}

func (m *MyLS) appendFile(str string) {
	if fInfo, err := os.Lstat(str); os.IsNotExist(err) {
		fmt.Printf("ls: cannot access '%s': No such file or directory\n", str)
	} else if os.IsPermission(err) {
		defer fmt.Printf("ls: cannot open directory '%s': Permission denied", str)
	} else if err == nil {
		if fInfo.IsDir() {
			m.Dirs = append(m.Dirs, str)
		} else {
			m.Files = append(m.Files, str)
		}
	}
}

func isFlag(str string) bool {
	if str[0] == '-' && len(str) > 1 {
		return true
	}
	return false
}
