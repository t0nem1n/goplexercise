package exer8_8

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirSize(dirs []string) {
	fileSize := make(chan int64)
	for _, dir := range dirs {
		go func(dir string) {
			walkDir(dir, fileSize)
			close(fileSize)
		}(dir)
	}

	var totalSize int64
	var totalFile int64
	for sz := range fileSize {
		totalSize += sz
		totalFile++
	}
	printDirSize(totalFile, totalSize)
}

func printDirSize(totalFile int64, totalSize int64) {
	fmt.Printf("%d file %.1f GB \n", totalFile, float64(totalSize)/(1<<30))
}

func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirEnts(dir) {
		if entry.IsDir() {
			subPath := filepath.Join(dir, entry.Name())
			walkDir(subPath, fileSize)
		} else {
			fileInfo, _ := entry.Info()
			fileSize <- fileInfo.Size()
		}
	}
}

func dirEnts(dir string) []os.DirEntry {
	if es, err := os.ReadDir(dir); err == nil {
		return es
	}
	return nil
}
