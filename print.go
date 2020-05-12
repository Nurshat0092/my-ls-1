package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"syscall"
)

func printLong(arr []os.FileInfo, pathStr *string, sizeLen, linkLen, totalM, giLen, uiLen int) {
	for _, w := range arr {
		s, _ := w.Sys().(*syscall.Stat_t)
		gID, uID, fSize, nLink := fmt.Sprint(s.Gid), fmt.Sprint(s.Uid), fmt.Sprint(s.Size), fmt.Sprint(s.Nlink)
		u, _ := user.LookupId(uID)
		g, _ := user.LookupGroupId(gID)

		fmt.Printf("%s ", strings.ToLower(w.Mode().String()))

		printSpaces(linkLen - len(nLink))
		fmt.Printf("%s ", nLink)

		printSpaces(uiLen - len(u.Username))
		fmt.Printf("%s ", u.Username)

		printSpaces(giLen - len(g.Name))
		fmt.Printf("%s ", g.Name)

		printSpaces(sizeLen - len(fSize))
		fmt.Printf("%s %s ", fSize, w.ModTime().Month().String()[:3])

		if dd := w.ModTime().Day(); dd < 10 {
			fmt.Printf(" %d ", dd)
		} else {
			fmt.Printf("%d ", dd)
		}

		if yy := w.ModTime().Year(); yy != currentYear {
			fmt.Printf(" %d ", yy)
		} else {
			fmt.Printf("%s ", w.ModTime().Format("15:04"))
		}

		if w.Mode()&os.ModeSymlink != 0 {
			dst, _ := os.Readlink(*pathStr + "/" + w.Name())
			_, err := os.Stat(dst)

			dstPath := *pathStr + "/" + dst
			if (len(m.Files) != 0) || (err == nil) {
				dstPath = dst
			}
			dstInfo, err := os.Stat(dstPath)
			if os.IsNotExist(err) {
				fmt.Printf(OrphanCol+" -> "+OrphanCol, w.Name(), dst)
			} else if os.IsPermission(err) {
				fmt.Print("***No Access***")
			} else if dstInfo.Mode()&os.ModeSymlink != 0 {
				fmt.Printf(SymLinkCol+" -> ", w.Name())
				fmt.Print(dst)
			} else {
				fmt.Printf(SymLinkCol+" -> ", w.Name())
				printName(dstInfo, dst)
			}
		} else {
			printName(w, w.Name())
		}
		fmt.Println()
	}
}

func printName(fInfo os.FileInfo, name string) {
	switch mode := fInfo.Mode(); {
	case mode.IsDir():
		isWritable, isSticky := mode.String()[len(mode.String())-2] == 'w', mode&os.ModeSticky != 0
		if isWritable && isSticky {
			fmt.Printf(StickyWritableCol, name)
		} else if !isWritable && isSticky {
			fmt.Printf(StickyCol, name)
		} else if isWritable && !isSticky {
			fmt.Printf(WritableCol, name)
		} else {
			fmt.Printf(DirCol, name)
		}
	case mode&os.ModeSymlink != 0:
		dst, _ := os.Readlink(name)
		_, err := os.Stat(dst)
		if os.IsNotExist(err) {
			fmt.Printf(OrphanCol, name)
		} else if os.IsPermission(err) {
			defer fmt.Printf("Permission Denied: %s\n", name)
		} else {
			fmt.Printf(SymLinkCol, name)
		}
	case mode&os.ModeNamedPipe != 0:
		fmt.Printf(PipeCol, name)
	case mode&os.ModeSocket != 0:
		fmt.Printf(SockCol, name)
	case mode&os.ModeDevice != 0:
		fmt.Printf(DeviceCol, name)
	case mode&os.ModeSetuid != 0:
		fmt.Printf(SetUIDCol, name)
	case mode&os.ModeSetgid != 0:
		fmt.Printf(SetGIDCol, name)
	default:
		if strings.Contains(mode.String(), "x") {
			fmt.Printf(ExecCol, name)
		} else {
			ext := getExtension(name)
			if archCompF[ext] {
				fmt.Printf(ArchCompCol, name)
			} else if audioF[ext] {
				fmt.Printf(AudioCol, name)
			} else if imageF[ext] {
				fmt.Printf(ImageCol, name)
			} else {
				fmt.Print(name)
			}
		}
	}
}

func getAllMax(arr []os.FileInfo) (int, int, int, int, int) {
	var blocks int64
	var maxSize int
	var linkLen int
	var giLen int
	var uiLen int

	for _, w := range arr {
		s, _ := w.Sys().(*syscall.Stat_t)
		lSize, gID, uID := len(fmt.Sprint(s.Nlink)), fmt.Sprint(s.Gid), fmt.Sprint(s.Uid)
		u, _ := user.LookupId(uID)
		g, _ := user.LookupGroupId(gID)
		blocks += s.Blocks
		if len(fmt.Sprint(s.Size)) > maxSize {
			maxSize = len(fmt.Sprint(s.Size))
		}
		if lSize > linkLen {
			linkLen = lSize
		}
		if len(u.Username) > uiLen {
			uiLen = len(u.Username)
		}
		if len(g.Name) > giLen {
			giLen = len(g.Name)
		}
	}
	blocks = (blocks + 1) / 2
	return maxSize, linkLen, int(blocks), giLen, uiLen
}

func removeHidden(arr []os.FileInfo) []os.FileInfo {
	for i := 0; i < len(arr); i++ {
		if arr[i].Name()[0] == '.' {
			arr[i] = arr[len(arr)-1]
			arr = arr[:(len(arr) - 1)]
			i--
		}
	}
	return arr
}

func printUsual(arr []os.FileInfo) {
	if len(arr) > 0 {
		for i := 0; i < len(arr)-1; i++ {
			printName(arr[i], arr[i].Name())
			fmt.Print("  ")
		}
		printName(arr[len(arr)-1], (arr[len(arr)-1]).Name())
		fmt.Println()
	}
}

func printWrittenFiles() {
	if m.Fl {
		subFiles := []os.FileInfo{}
		for _, w := range m.Files {
			fInfo, _ := os.Lstat(w)
			subFiles = append(subFiles, fInfo)
		}
		sizeLen, linkLen, totalM, giLen, uiLen := getAllMax(subFiles)
		for i := range subFiles {
			filePath := m.Files[i][:(len(m.Files[i])-len(subFiles[i].Name()))] + "."
			printLong([]os.FileInfo{subFiles[i]}, &filePath, sizeLen, linkLen, totalM, giLen, uiLen)
		}
	} else {
		for i, w := range m.Files {
			fInfo, _ := os.Lstat(w)
			printName(fInfo, w)
			if i != len(m.Files)-1 {
				fmt.Print("  ")
			}
		}
		if len(m.Files) > 0 {
			fmt.Println()
		}
	}
}

func printSpaces(count int) {
	for i := 0; i < count; i++ {
		fmt.Print(" ")
	}
}
