package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
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
	fmt.Println("fileName:", fileName)
	return nil
}
