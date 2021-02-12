package main

import (
	"errors"
	"fmt"
	nsema "gitee.com/taojun319/tjtools/control"
	_ "godev/内存/funcs"
	"godev/内存/nxlogdiy"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//var FSL = new(sync.RWMutex)

type SelectFiles struct {
	sync.RWMutex
	Files []string
	Path  string
}

var SelectFilesFree = sync.Pool{
	New: func() interface{} {
		return &SelectFiles{}
	},
}

func NewSelectFile() *SelectFiles {
	return &SelectFiles{}
}

//var files []string = make([]string, 0, 1<<13) // 不超过8192,可以重复利用
//var syncFile *sync.Mutex = new(sync.Mutex)

func (s *SelectFiles) GetFiles() []string {
	s.RLock()
	defer s.RUnlock()

	return s.Files
}

func (s *SelectFiles) Cleanfiles() {
	s.Lock()
	defer s.Unlock()
	s.Files = nil
}
func (s *SelectFiles) Len() uint64 {
	return uint64(len(s.Files))
}

func (s *SelectFiles) GetFileList() (err error) {
	f, err := os.Stat(s.Path)
	if err != nil {
		err = errors.New(fmt.Sprintf("path: %s ,Err:%s", s.Path, err))
		return err
	}
	if !f.IsDir() {
		err = errors.New(fmt.Sprintf("path: %s ,Not Dir", s.Path))
		return err
	}
	s.Lock()
	defer s.Unlock()
	s.Files = make([]string, 0)
	err = filepath.Walk(s.Path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		s.Files = append(s.Files, path)
		return nil
	})
	if err != nil {
		err = errors.New(fmt.Sprintf("path: %s ,err:%s", s.Path, err))
		return err
	}

	return nil
}

var FilesChan4 chan *SelectFiles = make(chan *SelectFiles, 0)
var sema4 *nsema.Semaphore = nsema.NewSemaphore(2)

func main() {

	nxlogdiy.InitLog()
	go Revice4()
	go Push4()

	select {}
}

func Push4() {

	for {
		rr := NewSelectFile()
		rr.Path = "/home/go/src/godev/内存"
		_ = rr.GetFileList()
		FilesChan4 <- rr
		time.Sleep(10 * time.Second)
	}

}

func Revice4() {
	for {
		select {
		case sa4 := <-FilesChan4:
			sema4.Acquire()
			nxlogdiy.GetLog().Info("File length:%d\n", len(sa4.GetFiles()))
			sema4.Release()
		}
	}
}
