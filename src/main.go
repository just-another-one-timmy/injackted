package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"structs"
)

const WORD_SEPS = " \n\t\\.*"

func addKeywords(index *structs.Index, doc string, keywords []string) {
	for _, keyword := range keywords {
		keyword = strings.Trim(keyword, WORD_SEPS)
		index.ConnectKeywordDoc(keyword, doc)
		fmt.Printf("Connected %v and %v\n", keyword, doc)
	}
}

func getKeywordsFromFile(fileName string) []string {
	fmt.Printf("%v\n", fileName)
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return nil
	}
	fmt.Printf("%v\n", string(fileContents))
	return strings.Split(string(fileContents), " ")
}

func BuildIndex() (*structs.Index, error) {
	result := structs.NewIndex()
	const base_dir = "./test_inputs/"

	dir, err := os.Open(base_dir)
	if err != nil {
		return nil, err
	}

	names, err := dir.Readdirnames(0)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return nil, err
	}

	rexp, err := regexp.Compile(".*-test")
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return nil, err
	}

	for _, name := range names {
		fmt.Printf("Found file %v\n", name)
		if rexp.MatchString(name) {
			addKeywords(result, base_dir+name, getKeywordsFromFile(base_dir+name))
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
