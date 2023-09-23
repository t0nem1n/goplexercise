package exer8_8

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type dirResult struct {
	dir  string
	size int64
	file int64
}

func (dr dirResult) String() string {
	return fmt.Sprintf("%s %d file %.1f GB", dr.dir, dr.file, float64(dr.size)/(1<<30))
}

type dirWalker struct {
	dir      string
	fileSize chan int64
	n        sync.WaitGroup
}

func newDirWalker(dir string) *dirWalker {
	return &dirWalker{
		dir,
		make(chan int64),
		sync.WaitGroup{},
	}
}

func (dw *dirWalker) walk() dirResult {
	dw.n.Add(1)
	go walkDir(dw.dir, &dw.n, dw.fileSize)
	go func() {
		dw.n.Wait()
		close(dw.fileSize)
	}()

	var totalSize int64
	var totalFile int64
	for sz := range dw.fileSize {
		totalSize += sz
		totalFile++
	}
	return dirResult{dw.dir, totalSize, totalFile}
}

func DirSize(dirs []string) {
	res := make(chan dirResult)
	var wg sync.WaitGroup
	for _, dir := range dirs {
		dw := newDirWalker(dir)
		wg.Add(1)
		go func() {
			res <- dw.walk()
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for dirRes := range res {
		fmt.Println(dirRes)
	}
}

func walkDir(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done()
	for _, entry := range dirEnts(dir) {
		if entry.IsDir() {
			subPath := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subPath, n, fileSize)
		} else {
			fileInfo, err := entry.Info()
			if err == nil {
				fileSize <- fileInfo.Size()
			}
		}
	}
}

var token = make(chan struct{}, 5)

func dirEnts(dir string) []os.DirEntry {
	token <- struct{}{}
	defer func() {
		<-token
	}()
	if es, err := os.ReadDir(dir); err == nil {
		return es
	}
	return nil
}
