package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/dereulenspiegel/makeup/makeup"
)

var inFile = flag.String("in", "", "Input file")
var fileType = flag.String("type", "", "Type of input (json)")

func main() {
	flag.Parse()
	if *inFile == "" {
		flag.Usage()
		return
	}
	var inType string
	if *fileType == "" {
		lastDot := strings.LastIndex(*inFile, ".")
		fileName := *inFile
		inType = fileName[lastDot+1:]
	}
	cosmetics := makeup.GetCosmetics(inType)
	if cosmetics != nil {
		out, err := cosmetics.PrettifyFile(*inFile)
		if err != nil {
			fmt.Printf("Error prettyfying file %s: %v", *inFile, err)
		} else {
			fmt.Printf("%s", string(out))
		}
	} else {
		fmt.Printf("Unknown format %s", inType)
	}
}
