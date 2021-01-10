package cmd

import (
	"convert-json-to-env/util"
	"flag"
	"fmt"
	"os"
)

var (
	file       string
	validity   bool
	fileFormat string
)

func Run() {
	printFlag := flag.Bool("print", false, "Print the conversion instead of outputting to a file.")
	flag.Parse()
	fileArgs := flag.Args()

	if len(fileArgs) == 1 {
		file = fileArgs[0]
		validity, fileFormat = util.CheckFileFormat(file)
	} else if len(fileArgs) > 1 || len(fileArgs) < 1 {
		fmt.Println("invalid value for file")
		os.Exit(0)
	}

	if validity != true {
		fmt.Printf("Invalid file format for file: %s. Need .json", fileFormat)
		os.Exit(0)
	}

	if fileFormat == "json" {
		ConvertJSON(file, *printFlag)
	} else {
		fmt.Println(`¯\_(ツ)_/¯`)
		os.Exit(0)
	}
}
