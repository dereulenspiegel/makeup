package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/dereulenspiegel/makeup/makeup"
)

var inFile = flag.String("in", "", "Input file")
var fileType = flag.String("type", "", "Type of input (json)")

func determineDataType(in []byte) string {
	for _, char := range in {
		if char == '{' || char == '[' {
			return "json"
		}
		if char == '<' {
			return "xml"
		}
	}
	return ""
}

func main() {
	flag.Parse()
	var uglyData []byte
	var prettyData []byte
	var err error
	if *inFile != "" {
		uglyData, err = ioutil.ReadFile(*inFile)
		if err != nil {
			log.Fatalf("Can't read file %s: %v", *inFile, err)
		}
	} else {
		uglyData, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Can't read from stdin: %v", err)
		}
	}
	var inType string
	if *fileType == "" && *inFile != "" {
		lastDot := strings.LastIndex(*inFile, ".")
		fileName := *inFile
		inType = fileName[lastDot+1:]
	} else if *fileType != "" {
		inType = *fileType
	} else {
		inType = determineDataType(uglyData[:32])
	}
	cosmetics := makeup.GetCosmetics(inType)

	if cosmetics != nil {
		prettyData, err = cosmetics.Prettify(uglyData)
		if err != nil {
			log.Fatalf("Can't prettify input data: %v", err)
		}
		os.Stdout.Write(prettyData)
	} else {
		log.Fatalf("Can't determine type of input data")
	}
}
