package main

import (
	"log"
	"os"
	"github.com/r0nni3/text-replace"
)

func main() {
	inputFile, dictFile := replacer.GetInputs(os.Args)

	dict, err := replacer.LoadDictionary(dictFile, "")
	if err != nil {
		log.Print("Error dictionary: ")
		log.Fatalln(err.Error())
	}

	err = replacer.ReplaceText(inputFile, dict)
	if err != nil {
		log.Print("Error replace: ")
		log.Panicln(err.Error())
	}

	log.Println("All Replaces Done!!")
}
