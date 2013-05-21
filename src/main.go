package main

import (
	"flag"
	"fmt"
	"structs"
)

func BuildIndex() *structs.Index {
	result := structs.NewIndex()
	return result
}

func main() {
	flag.Parse()

	index := BuildIndex()
	index.ConnectKeywordDoc("ABBA", 1)
	index.ConnectKeywordDoc("ABBA", 2)
	index.ConnectKeywordDoc("Beatles", 3)

	for _, query := range flag.Args() {
		docs := index.GetDocsByKeyword(query)
		fmt.Printf("Docs for keyword %v:\n", query)
		for doc, _ := range *docs.IterableMap() {
			fmt.Printf("\t%v\n", doc)
		}
	}
}
