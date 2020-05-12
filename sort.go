package main

import (
	"math/rand"
	"os"
	"strings"
)

func sortFiles(arr interface{}) {
	switch t := arr.(type) {
	case []string:
		quickSortStr(t)
	case []os.FileInfo:
		quickSort(t)
	}
	if m.Fr {
		reverseFiles(arr)
	}
}

func quickSortStr(files []string) {
	if len(files) < 2 {
		return
	}
	left, right := 0, len(files)-1
	pivot := rand.Int() % len(files)
	files[pivot], files[right] = files[right], files[pivot]
	switch m.Ft {
	case true:
		for i := range files {
			f1, _ := os.Lstat(files[i])
			f2, _ := os.Lstat(files[right])
			if f1.ModTime().UnixNano() > f2.ModTime().UnixNano() {
				files[left], files[i] = files[i], files[left]
				left++
			} else if f1.ModTime().UnixNano() == f2.ModTime().UnixNano() {
				if compareNames(f1, f2) {
					files[left], files[i] = files[i], files[left]
					left++
				}
			}
		}
	case false:
		for i := range files {
			f1, _ := os.Lstat(files[i])
			f2, _ := os.Lstat(files[right])
			if compareNames(f1, f2) {
				files[left], files[i] = files[i], files[left]
				left++
			}
		}
	}
	files[left], files[right] = files[right], files[left]
	quickSortStr(files[:left])
	quickSortStr(files[left+1:])
}

func quickSort(files []os.FileInfo) {
	if len(files) < 2 {
		return
	}
	left, right := 0, len(files)-1
	pivot := rand.Int() % len(files)
	files[pivot], files[right] = files[right], files[pivot]
	switch m.Ft {
	case true:
		for i := range files {
			if files[i].ModTime().UnixNano() > files[right].ModTime().UnixNano() {
				files[left], files[i] = files[i], files[left]
				left++
			} else if files[i].ModTime().UnixNano() == files[right].ModTime().UnixNano() {
				if compareNames(files[i], files[right]) {
					files[left], files[i] = files[i], files[left]
					left++
				}
			}
		}
	case false:
		for i := range files {
			if compareNames(files[i], files[right]) {
				files[left], files[i] = files[i], files[left]
				left++
			}
		}
	}
	files[left], files[right] = files[right], files[left]
	quickSort(files[:left])
	quickSort(files[left+1:])
}

func compareNames(i, right os.FileInfo) bool {
	iName, rName := strings.ToLower(i.Name()), strings.ToLower(right.Name())
	if iName[0] == '.' {
		iName = iName[1:]
	}
	if rName[0] == '.' {
		rName = rName[1:]
	}
	if (iName < rName) || ((iName == rName) && (i.Name() > right.Name())) {
		return true
	}
	return false
}

func reverseFiles(files interface{}) {
	switch t := files.(type) {
	case []string:
		for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
			t[i], t[j] = t[j], t[i]
		}
	case []os.FileInfo:
		for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
			t[i], t[j] = t[j], t[i]
		}
	}
}
