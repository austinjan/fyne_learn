package main

import (
	"os"
	"path"
	"runtime"
)

// checkFiles  checks if the files exist

func checkFiles(loc string) bool {
	files := []string{
		"bbrootsvc",
		"bbanomsvc",
		"bbidpsvc",
		"bblogsvc",
		"bbnimbl",
	}
	exist := true
	for _, file := range files {
		if runtime.GOOS == "windows" {
			file = file + ".exe"
		}
		// check file exist
		fullpath := path.Join(loc, file)
		// check fullpath exist
		_, err := os.Stat(fullpath)
		if err != nil || os.IsNotExist(err) {
			exist = false
			break
		}
	}
	return exist
}
