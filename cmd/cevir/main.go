package main

import (
	"flag"
	"log"
	"os"

	"github.com/maliByatzes/cevir"
)

func main() {
	// TODO: provide more robust flags
	var fileName string
	flag.StringVar(&fileName, "file", "", "File to convert")

	flag.Parse()

	if fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := runApp(fileName); err != nil {
		log.Fatal(err)
	}
}

func runApp(fileName string) error {
	// call convert epub to pdf for now
	if err := cevir.ConvertEpubToPDF(fileName); err != nil {
		return err
	}
	return nil
}
