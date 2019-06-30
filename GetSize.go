package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	_  = iota
	kb = 1 << (10 * iota)
	mb
	gb
	tb
)

func getFileSize(checkFile int64) string {
	sf := float32(checkFile)
	commas := float32(len(strconv.Itoa(int(checkFile)))) / 4.00
	if commas <= 1.00 {
		return fmt.Sprintf("%.2fBytes\n", sf)
	} else if commas > 1.00 && commas <= 2.00 {
		return fmt.Sprintf("%.2fKB\n", sf/kb)
	} else if commas > 2.00 && commas <= 3.00 {
		return fmt.Sprintf("%.2fMB\n", sf/mb)
	} else if commas > 4.00 && commas <= 5.00 {
		return fmt.Sprintf("%.2fGB\n", sf/gb)
	} else if commas > 5.00 && commas <= 6.00 {
		return fmt.Sprintf("%.2fTB\n", sf/tb)
	}
	return "nil"
}

func extractor(checkFile string) int64 {
	info, err1 := os.Stat(checkFile)
	if err1 != nil {
		log.Fatal(err1)
	}
	return info.Size()
}

// FileSize - Get the file size in a human readable form
func FileSize(checkFile string) string {
	return getFileSize(extractor(checkFile))
}

func main() {
	checkFile := flag.String("filename", "", "Full Path of file")
	flag.Parse()
	fmt.Printf(getFileSize(extractor(*checkFile)))
}
