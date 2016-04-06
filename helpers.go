package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type dicEntry struct {
	PrevEntry string
	NewEntry  string
}

// DefaultSeparator default character separator for the CSV file
const DefaultSeparator = '\t'

func loadDictionary(dictFile string, separator string) ([]*dicEntry, error) {
	file, err := os.Open(dictFile)
	if err != nil {
		return nil, err
	}

	var dict = make([]*dicEntry, 0)
	in := csv.NewReader(bufio.NewReader(file))
	in.Comma = DefaultSeparator
	if separator != "" {
		in.Comma = rune(separator[0])
	}

	for {
		record, err := in.Read()
		if err == io.EOF {
			break
		}

		entry := new(dicEntry)
		entry.PrevEntry = strings.TrimSpace(record[0])
		entry.NewEntry = strings.TrimSpace(record[1])

		dict = append(dict, entry)
	}

	return dict, nil
}

func replaceText(document string, dict []*dicEntry) error {
	read, err := ioutil.ReadFile(document)
	if err != nil {
		return err
	}

	newcontent := string(read)
	for i := range dict {
		newcontent = strings.Replace(
			newcontent,
			dict[i].PrevEntry,
			dict[i].NewEntry, -1)
	}

	err = ioutil.WriteFile(document, []byte(newcontent), 0)
	if err != nil {
		return err
	}

	return nil
}

func getInputs(args []string) (string, string) {
	if len(args[1:]) < 2 {
		log.Fatalln("replaceScript <INPUT_FILE> <DICTIONARY>")
	}

	inputFile := args[1]
	dictFile := args[2]

	return inputFile, dictFile
}
