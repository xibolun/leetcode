package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getTarget(name string) string {
	if len(name) <= 20 {
		return name
	}

	ret, err := strconv.Atoi(name[len(name)-8:])

	if err != nil {
		fmt.Printf("fail convert value to int, %s", name[len(name)-8:0])
		return name
	}
	return strings.Replace(name, name[len(name)-8:], strconv.Itoa(ret+1340), -1)
}

func loopDir(dir string, fsPath []string) []string {
	file, _ := os.Stat(dir)
	fsPath = append(fsPath, filepath.Base(dir))
	if !file.IsDir() {
		return fsPath
	}
	fsEntries, _ := os.ReadDir(dir)
	for _, item := range fsEntries {
		if item.IsDir() {
			loopDir(filepath.Join(dir, item.Name()), fsPath)
		}
		fsPath = append(fsPath, filepath.Join(dir, item.Name()))
	}
	return fsPath
}

var srcPath = []string{"/tmp/test"}

func main() {
	for _, item := range srcPath {
		allFS := make([]string, 0)
		allFS = loopDir(item, allFS)

		for _, fs := range allFS {
			if err := os.Rename(fs, getTarget(fs)); err != nil {
				fmt.Printf("fail rename file %s, %s", fs, err.Error())
			}
		}
	}
}
