package main

// The du2 variant uses select and a time.Ticker
// to print the totals periodically if -v is set.

import (
	// "flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//!+
// var verbose = flag.Bool("v", true, "show verbose progress messages")
type fs struct {
	nf int
	fs int64
}

// var fss chan []fs

func main() {
	// ...start background goroutine...

	//!-
	// Determine the initial directories.
	// flag.Parse()
	// roots := flag.Args()

	// if len(roots) == 0 {
	dir := "c:/godev"
	roots, _ := ioutil.ReadDir(dir)
	// roots := []string{"c:/godev"}
	// }

	// Traverse the file tree.
	// n := 0
	for _, r := range roots {
		// fmt.Println(r.Name())
		// n++

		// for _, root := range r {
		// fmt.Println(r.Name())
		// fmt.Printf("%s,%t\n", r.Name(), r.IsDir())
		if r.IsDir() {
			go func(dd string) {
				// var nfiles int
				// var fileSizes int64
				// dd := r.Name()
				var fss fs

				walkDir(filepath.Join(dir, dd), &fss)
				fmt.Printf("dir:%s,size:%.1f B,nf:%d\n", dd, float64(fss.fs), fss.nf)

			}(r.Name())
		} else {
			// fileSizes += r.Size()
			fmt.Printf("file:%s,size:%d B\n", r.Name(), r.Size())
		}

		// printDiskUsage(r.Name(), nf, fileSizes)

	}
	// fmt.Println(n)
	time.Sleep(10 * time.Second)
	//!+
	// Print the results periodically.
	// var tick <-chan time.Time
	// if *verbose {
	// tick = time.Tick(10 * time.Millisecond)
	// }

	// loop:
	// 	for {

	// 		select {
	// 		case size, ok := <-fileSizes:
	// 			if !ok {
	// 				break loop // fileSizes was closed
	// 			}
	// 			nfiles++
	// 			nbytes += size
	// 		case <-tick:
	// 			printDiskUsage(nfiles, nbytes)
	// 		}
	// 	}
	// 	printDiskUsage(nfiles, nbytes) // final totals
	// }
}

//!-

// func printDiskUsage(r string, nf int, nbytes int64) {
// 	fmt.Printf("dir:%s,%d files  %.1f KB\n", r, nf, float64(nbytes)/1024)
// }

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fss *fs) {

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fss)

		} else {
			fss.fs += entry.Size()
			fss.nf += 1

		}

	}
	// return nfiles, fileSizes
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	// fmt.Println(dir)
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
