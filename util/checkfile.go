package util

import (
	"fmt"
	"os"
	"strings"
)

var AcceptedFormats = []string{"json", "env"}

func CheckFileFormat(filename string) (bool, string) {
	for _, i := range AcceptedFormats {
		if strings.HasSuffix(string(filename), i) {
			return true, i
		}
	}
	return false, filename
}

func DoesFileExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("issue closing file")
		os.Exit(0)
	}
}
