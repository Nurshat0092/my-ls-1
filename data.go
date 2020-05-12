package main

import (
	"os"
	"time"
)

// const FileCol = "\033[00m%s\033[0m"
const DirCol = "\033[01;34m%s\033[0m"
const SymLinkCol = "\033[01;36m%s\033[0m"
const PipeCol = "\033[40;33m%s\033[0m"
const SockCol = "\033[01;35m%s\033[0m"
const DeviceCol = "\033[40;33;01m%s\033[0m"
const OrphanCol = "\033[40;31;01m%s\033[0m"
const SetUIDCol = "\033[37;41m%s\033[0m"
const SetGIDCol = "\033[30;43m%s\033[0m"
const CapCol = "\033[30;41m%s\033[0m"
const StickyWritableCol = "\033[30;42m%s\033[0m"
const WritableCol = "\033[34;42m%s\033[0m"
const StickyCol = "\033[37;44m%s\033[0m"
const ExecCol = "\033[01;32m%s\033[0m"
const AudioCol = "\033[00;36m%s\033[0m"
const ImageCol = "\033[01;35m%s\033[0m"
const ArchCompCol = "\033[01;31m%s\033[0m"

type MyLS struct {
	Fa    bool
	FR    bool
	Fl    bool
	Fr    bool
	Ft    bool
	Files []string
	Dirs  []string
}

var (
	m           *MyLS
	manyDirs    bool
	currentYear int
	audioF      map[string]bool
	imageF      map[string]bool
	archCompF   map[string]bool
)

func init() {
	m = &MyLS{}
	args := os.Args[1:]
	m.readArgs(args)
	currentYear = time.Now().Year()
	audioF = map[string]bool{
		".aac": true, ".au": true, ".flac": true, ".ogg": true, ".ra": true,
		".m4a": true, ".mid": true, ".midi": true, ".wav": true, ".oga": true,
		".mka": true, ".mp3": true, ".mpc": true, ".opus": true, ".spx": true, ".xspf": true,
	}
	imageF = map[string]bool{
		".jpg": true, ".mjpeg": true, ".pbm": true, ".tga": true, ".tif": true, ".svg": true,
		".jpeg": true, ".gif": true, ".pgm": true, ".xbm": true, ".tiff": true, ".svgz": true,
		".mjpg": true, ".bmp": true, ".ppm": true, ".xpm": true, ".png": true, ".mng": true,
		".pcx": true, ".mov": true, ".mpg": true, ".mpeg": true, ".m2v": true, ".mkv": true,
		".webm": true, ".ogm": true, ".mp4": true, ".m4v": true, ".mp4v": true, ".vob": true,
		".qt": true, ".nuv": true, ".wmv": true, ".asf": true, ".rm": true, ".rmvb": true,
		".flc": true, ".avi": true, ".fli": true, ".flv": true, ".gl": true, ".dl": true,
		".xcf": true, ".xwd": true, ".yuv": true, ".cgm": true, ".emf": true, ".ogv": true, ".ogx": true,
	}
	archCompF = map[string]bool{
		".tar": true, ".tgz": true, ".arc": true, ".arj": true, ".taz": true, ".lha": true,
		".lz4": true, ".lzh": true, ".lzma": true, ".tlz": true, ".txz": true, ".tzo": true,
		".t7z": true, ".zip": true, ".z": true, ".Z": true, ".dz": true, ".gz": true,
		".lrz": true, ".lz": true, ".lzo": true, ".xz": true, ".zst": true, ".tzst": true,
		".bz2": true, ".bz": true, ".tbz": true, ".tbz2": true, ".tz": true, ".deb": true,
		".rpm": true, ".ear": true, ".alz": true, ".cpio": true, ".cab": true, ".dwm": true,
		".jar": true, ".sar": true, ".ace": true, ".7z": true, ".wim": true, ".esd": true,
		".war": true, ".rar": true, ".zoo": true, ".rz": true, ".swm": true,
	}
}

func getExtension(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

/*
for x in {0..8}; do for i in {30..37}; do for a in {40..47}; do echo -ne "\e[$x;$i;$a""m\\\e[$x;$i;$a""m\e[0;37;40m "; done; echo; done; done; echo ""
*/
/*
dircolors -p
*/
