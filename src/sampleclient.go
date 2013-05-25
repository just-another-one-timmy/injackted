package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"structs"
	"regexp"
)

func loadIndex(index *structs.Index, fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	gobIndex := &structs.GobIndex{nil, nil}

	decoder := gob.NewDecoder(f)
	if err := decoder.Decode(gobIndex); err != nil {
		return err
	}

	*index = *structs.GobIndexToIndex(gobIndex)
	return nil
}

func saveIndex(index *structs.Index, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(structs.IndexToGobIndex(index))

	return err
}

func connect(index *structs.Index, keyword string, doc string) {
	index.ConnectKeywordDoc(keyword, doc)
	fmt.Printf("%q and %q were connected in the index\n", keyword, doc)
}

func listDocs(index *structs.Index, keyword string) {
	fmt.Printf("Documents which contain keyword %q:\n", keyword)
	count := 0
	for doc := range index.IteratorDocsByKeyword(keyword) {
		count++
		fmt.Printf("\t%v\n", doc)
	}
	fmt.Printf("%v in total\n", count)
}

func listKeywords(index *structs.Index, doc string) {
	fmt.Printf("Keywords contained in %q:\n", doc)
	count := 0
	for keyword := range index.IteratorKeywordsByDoc(doc) {
		count++
		fmt.Printf("\t%v\n", keyword)
	}
	fmt.Printf("%v in total\n", count)
}

func listAllKeywords(index *structs.Index) {
	fmt.Println("There are next keywords in the index:")
	count := 0
	for keyword := range index.IteratorKeywords() {
		count++
		fmt.Printf("\t%v\n", keyword)
	}
	fmt.Printf("%v in total\n", count)
}

func listAllDocs(index *structs.Index) {
	fmt.Println("There are next docs in the index:")
	count := 0
	for doc := range index.IteratorDocs() {
		count++
		fmt.Printf("\t%v\n", doc)
	}
	fmt.Printf("%v in total\n", count)
}

func listAll(index *structs.Index) {
	fmt.Println("Contents of the index:")
	for doc := range index.IteratorDocs() {
		fmt.Printf("%v\n", doc)
		for keyword := range index.IteratorKeywordsByDoc(doc) {
			fmt.Printf("\t%v\n", keyword)
		}
	}
}

func queryIndex(index *structs.Index, rexpstr string) {
	r, err := regexp.Compile(rexpstr);
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}
	matchedDocsCount, matchedKeywordsCount := 0, 0
	fmt.Println("Docs matched by regexp:")
	for doc := range index.IteratorDocs() {
		docStr := doc.(string)
		if r.MatchString(docStr) {
			matchedDocsCount++
			fmt.Println(docStr)
		}
	}
	fmt.Printf("%v total.\n", matchedDocsCount)
	fmt.Println("Keywords matched by regexp:")
	for keyword := range index.IteratorKeywords() {
		keywordStr := keyword.(string)
		if r.MatchString(keywordStr) {
			fmt.Println(keyword)
			matchedKeywordsCount++
			inDocs := make([]string, 0)
			for doc := range index.IteratorDocsByKeyword(keywordStr) {
				inDocs = append(inDocs, doc.(string))
			}
			fmt.Printf("\tFound in the next %v documents:\n", len(inDocs))
			for _, doc := range inDocs {
				fmt.Printf("\t%v\n", doc)
			}
		}
	}
}

// Returns true if quit command received.
func handleCommand(command string, index *structs.Index) bool {
	switch command {
	case "help":
		fmt.Printf("Available commands:\n" +
			"help:                       displays help\n" +
			"load <filename>:            loads index from <filename>\n" +
			"q <regexp-query>:           searches the stuff\n" +
			"save <filename>:            saves index to <filename>\n" +
			"connect <keyword> <doc>:    connects <keyword> and <doc> in index\n" +
			"list-docs <keyword>:        lists documents containing <keyword>\n" +
			"list-keywords <doc>:        lists keywords contained in <doc>\n" +
			"list-all-keywords:          lists all keywords\n" +
			"list-all-docs:              lists all docs\n" +
			"list-all:                   lists all docs and keywords\n" +
			"bye-bye:                    quits\n")
	case "load":
		var arg string
		fmt.Scan(&arg)
		if err := loadIndex(index, arg); err != nil {
			fmt.Printf("Error while loading index: %v\n", err.Error())
		} else {
			fmt.Println("Successfully loaded index.")
		}
		return false
	case "q":
		var arg string
		fmt.Scan(&arg)
		queryIndex(index, arg)
		return false
	case "save":
		var arg string
		fmt.Scan(&arg)
		if err := saveIndex(index, arg); err != nil {
			fmt.Printf("Error while saving index: %v\n", err.Error())
		} else {
			fmt.Println("Successfully saved index.")
		}
		return false
	case "connect":
		var arg1, arg2 string
		fmt.Scan(&arg1, &arg2)
		connect(index, arg1, arg2)
		return false
	case "list-docs":
		var arg string
		fmt.Scan(&arg)
		listDocs(index, arg)
		return false
	case "list-keywords":
		var arg string
		fmt.Scan(&arg)
		listKeywords(index, arg)
		return false
	case "list-all-keywords":
		listAllKeywords(index)
		return false
	case "list-all-docs":
		listAllDocs(index)
		return false
	case "bye-bye":
		fmt.Printf("bye-bye too!\n")
		return true
	case "list-all":
		listAll(index)
		return false
	default:
		fmt.Printf("Unrecognized command: %q\n", command)
		return false
	}

	// To make compiler happy.
	return false
}

func main() {
	index := structs.NewIndex()
	for {
		fmt.Println("Enter command (e.g. 'help') and arguments.")
		var userInput string
		_, err := fmt.Scan(&userInput)
		if err != nil {
			if err == io.EOF {
				return
			} else {
				fmt.Println(err.Error())
			}
		} else {
			if handleCommand(userInput, index) {
				return
			}
		}
	}
}
