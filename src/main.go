package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"structs"
)

func addKeywords(index *structs.Index, doc string, keywords []string) {
	for _, keyword := range keywords {
		index.ConnectKeywordDoc(keyword, doc)
	}
}

func getKeywordsFromFile(file string) []string {
	// TODO(iaroslav): write actual code.
	return []string{"one", "two", "three", "four", "five"}
}

func BuildIndex() (*structs.Index, error) {
	result := structs.NewIndex()

	dir, err := os.Open(".")
	if err != nil {
		return nil, err
	}

	names, err := dir.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	rexp, err := regexp.Compile(".*-test")
	if err != nil {
		return nil, err
	}

	for _, name := range names {
		if rexp.MatchString(name) {
			addKeywords(result, name, getKeywordsFromFile(name))
		}
	}

	return result, nil
}

func main() {
	flag.Parse()

	index, err := BuildIndex()

	if err != nil {
		fmt.Printf("Error while building index: %v.\n", err.Error())
		return
	}

	for _, query := range flag.Args() {
		docs := index.GetDocsByKeyword(query)
		fmt.Printf("Docs for keyword %v:\n", query)
		for doc, _ := range *docs.IterableMap() {
			fmt.Printf("\t%v\n", doc)
		}
	}
}
