package main

import (
	"fmt"
	"io"
	"structs"
)

func loadIndex(index *structs.Index, fileName string) {
	//TODO(iaroslav): write real code.
}

func saveIndex(index *structs.Index, fileName string) {
	//TODO(iaroslav): write real code.
}

func connect(index *structs.Index, keyword string, doc string) {
	index.ConnectKeywordDoc(keyword, doc)
	fmt.Printf("%q and %q were connected in the index\n", keyword, doc)
}

func listDocs(index *structs.Index, keyword string) {
	fmt.Printf("Documents which contain keyword %q:\n", keyword)
	for doc, _ := range *index.GetDocsByKeyword(keyword).IterableMap() {
		fmt.Printf("\t%v\n", doc)
	}
}

func listKeywords(index *structs.Index, doc string) {
	fmt.Printf("Keywords contained in %q:\n", doc)
	for keyword, _ := range *index.GetKeywordsByDoc(doc).IterableMap() {
		fmt.Printf("\t%v\n", keyword)
	}
}

func handleCommand(command string, index *structs.Index) {
	switch command {
	case "help":
		fmt.Printf("Available commands:\n" +
			"help:                       displays help\n" +
			"load <filename>:            loads index from <filename>\n" +
			"save <filename>:            saves index to <filename>\n" +
			"connect <keyword> <doc>:    connects <keyword> and <doc> in index\n" +
			"list-docs <keyword>:        lists documents containing <keyword>\n" +
			"list-keywords <doc>:        lists keywords contained in <doc>\n")
	case "load":
		var arg string
		fmt.Scan(&arg)
		loadIndex(index, arg)
	case "save":
		var arg string
		fmt.Scan(&arg)
		saveIndex(index, arg)
	case "connect":
		var arg1, arg2 string
		fmt.Scan(&arg1, &arg2)
		connect(index, arg1, arg2)
	case "list-docs":
		var arg string
		fmt.Scan(&arg)
		listDocs(index, arg)
	case "list-keywords":
		var arg string
		fmt.Scan(&arg)
		listKeywords(index, arg)
	default:
		fmt.Printf("Unrecognized command: %q\n", command)
	}
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
			handleCommand(userInput, index)
		}
	}
}
