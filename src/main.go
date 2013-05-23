package main

import (
	"fmt"
	"io"
	"os"
	"structs"
)

func loadIndex(index *structs.Index, fileName string) {
	// TODO(iaroslav): clear index before loading.
	// For now, let it be.
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error while loading index.")
		fmt.Printf("%v\n", err.Error())
		return
	}
	defer func() {
		if err := inputFile.Close(); err != nil {
			fmt.Println("Error while loading index.")
			fmt.Printf("%v\n", err.Error())
		} else {
			fmt.Printf("Successfully loaded index from %v\n", fileName)
		}
	}()

	for {
		var keyword string
		var docsCount int
		if _, err := fmt.Fscan(inputFile, &keyword, &docsCount); err != nil {
			if err != io.EOF {
				fmt.Println("Error while loading index.")
				fmt.Printf("%v\n", err.Error())
			}
			return
		}
		for i := 0; i < docsCount; i++ {
			var doc string
			if _, err := fmt.Fscan(inputFile, &doc); err != nil {
				fmt.Println("Error while loading index.")
				fmt.Printf("%v\n", err.Error())
			}
			index.ConnectKeywordDoc(keyword, doc)
		}
	}
}

func saveIndex(index *structs.Index, fileName string) {
	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while saving index.")
		fmt.Printf("%v\n", err.Error())
		return
	}
	defer func() {
		if err := outputFile.Close(); err != nil {
			fmt.Println("Error while saving index.")
			fmt.Printf("%v\n", err.Error())
		} else {
			fmt.Printf("Successfully saved index to %v\n", fileName)
		}
	}()
	if _, err := index.Print(outputFile); err != nil {
		fmt.Printf("Error while savin index.")
		fmt.Printf("%v\n", err.Error())
	}
}

func connect(index *structs.Index, keyword string, doc string) {
	index.ConnectKeywordDoc(keyword, doc)
	fmt.Printf("%q and %q were connected in the index\n", keyword, doc)
}

func listDocs(index *structs.Index, keyword string) {
	fmt.Printf("Documents which contain keyword %q:\n", keyword)
	for doc := range index.IteratorDocsByKeyword(keyword) {
		fmt.Printf("\t%v\n", doc)
	}
}

func listKeywords(index *structs.Index, doc string) {
	fmt.Printf("Keywords contained in %q:\n", doc)
	for keyword := range index.IteratorKeywordsByDoc(doc) {
		fmt.Printf("\t%v\n", keyword)
	}
}

func listAllKeywords(index *structs.Index) {
	fmt.Println("There are next keywords in the index:")
	for keyword := range index.IteratorKeywords() {
		fmt.Printf("\t%v\n", keyword)
	}
}

func listAllDocs(index *structs.Index) {
	fmt.Println("There are next docs in the index:")
	for doc := range index.IteratorDocs() {
		fmt.Printf("\t%v\n", doc)
	}
}

// Returns true if quit command received.
func handleCommand(command string, index *structs.Index) bool {
	switch command {
	case "help":
		fmt.Printf("Available commands:\n" +
			"help:                       displays help\n" +
			"load <filename>:            loads index from <filename>\n" +
			"save <filename>:            saves index to <filename>\n" +
			"connect <keyword> <doc>:    connects <keyword> and <doc> in index\n" +
			"list-docs <keyword>:        lists documents containing <keyword>\n" +
			"list-keywords <doc>:        lists keywords contained in <doc>\n" +
			"list-all-keywords:          lists all keywords\n" +
			"list-all-docs:              lists all docs\n" +
			"bye-bye:                    quits\n")
	case "load":
		var arg string
		fmt.Scan(&arg)
		loadIndex(index, arg)
		return false
	case "save":
		var arg string
		fmt.Scan(&arg)
		saveIndex(index, arg)
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
