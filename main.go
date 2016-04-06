package main

import (
	"log"
	"os"
	"path/filepath"
)

var currentDir, fileExtension string

func main() {

	inputFile, dictFile := getInputs(os.Args)
	currentDir, _ = filepath.Abs(filepath.Dir(inputFile))
	fileExtension = filepath.Ext(inputFile)

	dict, err := loadDictionary(dictFile, "")
	if err != nil {
		log.Print("Error dictionary: ")
		log.Fatalln(err.Error())
	}

	err = replaceText(inputFile, dict)
	if err != nil {
		log.Print("Error replace: ")
		log.Panicln(err.Error())
	}

	log.Println("All Replaces Done!!")
}
