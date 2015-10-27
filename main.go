package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dereulenspiegel/makeup/makeup"
)

var inFile = flag.String("in", "", "Input file")
var fileType = flag.String("type", "", "Type of input (json)")

func main() {
	flag.Parse()
	var inType string
	if *fileType == "" {
		lastDot := strings.LastIndex(*inFile, ".")
		fileName := *inFile
		inType = fileName[lastDot+1:]
	} else {
		inType = *fileType
	}
	cosmetics := makeup.GetCosmetics(inType)

	if cosmetics != nil {
		var err error
		var out []byte
		if *inFile == "" {
			var in []byte
			in, err = ioutil.ReadAll(os.Stdin)
			out, err = cosmetics.Prettify(in)
		} else {
			out, err = cosmetics.PrettifyFile(*inFile)
		}
		if err != nil {
			fmt.Printf("Error prettyfying file %s: %v", *inFile, err)
		} else {
			fmt.Printf("%s", string(out))
		}
	} else {
		fmt.Printf("Unknown format %s", inType)
	}
}
