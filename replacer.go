// Package text-replacer implements utility routines for
// replacing text on a target file provided by a dictionary.
package replacer

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// DicEntry maps the string to subsitute to the replacement string
type DicEntry struct {
	PrevEntry string
	NewEntry  string
}

// DefaultSeparator default character separator for the CSV file
const DefaultSeparator = '\t'

// LoadDictionary loads the dictionary file to memory.
// Takes dictFile string and separator character as string
// It Returns a Pointer Array of DicEntries if loaded successfully or
// error if not.
func LoadDictionary(dictFile string, separator string) ([]*DicEntry, error) {
	file, err := os.Open(dictFile)
	if err != nil {
		return nil, err
	}

	var dict = make([]*DicEntry, 0)
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

		entry := new(DicEntry)
		entry.PrevEntry = strings.TrimSpace(record[0])
		entry.NewEntry = strings.TrimSpace(record[1])

		dict = append(dict, entry)
	}

	return dict, nil
}

// ReplaceText loads the target text file and replaces text from dicEntries array
// Takes parameters document file path string and array of DicEntry pointers
// It returns error if couldn't open file or if couldn't write to file.
func ReplaceText(document string, dict []*DicEntry) error {
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

// GetInputs parses command line parameters
func GetInputs(args []string) (string, string) {
	if len(args[1:]) < 2 {
		log.Fatalln("replaceScript <INPUT_FILE> <DICTIONARY>")
	}

	inputFile := args[1]
	dictFile := args[2]

	return inputFile, dictFile
}
