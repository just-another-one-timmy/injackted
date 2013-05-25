package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"structs"
)

type DocKeywordPair struct {
	Doc     string
	Keyword string
}

func sendTrue(w chan bool) {
	w <- true
}

func reportError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err.Error())
}

func processFile(fileName string, c chan DocKeywordPair, w chan bool) {
	//c <- *(&DocKeywordPair{fileName, "tempkeyword!"})
	// Inform receiver that current goroutine is done.
	defer sendTrue(w)

	if f, err := os.Open(fileName); err == nil {
		for {
			var keyword string
			_, err = fmt.Fscan(f, &keyword)
			if err != nil {
				if err != io.EOF {
					reportError(err)
				}
				break
			} else {
				c <- DocKeywordPair{fileName, keyword}
			}
		}
	}
}

func crawl(index *structs.Index, fileNames []string) {
	// channel to recieve doc-keyword pairs.
	c := make(chan DocKeywordPair)
	// channel to count number of go-routines.
	w := make(chan bool)
	goroutinesCount := len(fileNames)
	for _, fileName := range fileNames {
		go processFile(fileName, c, w)
	}
	for goroutinesCount > 0 {
		select {
		case docKeywordPair := <-c:
			index.ConnectKeywordDoc(docKeywordPair.Keyword,
				docKeywordPair.Doc)
		case _ = <-w:
			goroutinesCount--
		}
	}

	// I am wondering, might it be the case that goroutine
	// had sent values to c and w, and value from w was read first?
	// Then c still contains the last keyword from the file.
	select {
	case docKeywordPair := <-c:
		index.ConnectKeywordDoc(docKeywordPair.Keyword,
			docKeywordPair.Doc)
	default:
		// nothing to do here, but we're sure that
		// nothing is in the c channel.
	}
}

// Copy-paste-hacked from sampleclient.go.
func saveIndex(index *structs.Index) error {
	encoder := gob.NewEncoder(os.Stdout)
	err := encoder.Encode(structs.IndexToGobIndex(index))
	return err
}

func main() {
	fileNames := make([]string, 0)
	for {
		var fileName string
		_, err := fmt.Scan(&fileName)
		if err != nil {
			if err != io.EOF {
				reportError(err)
			}
			break
		}
		fileNames = append(fileNames, fileName)
	}

	index := structs.NewIndex()
	crawl(index, fileNames)

	if err := saveIndex(index); err != nil {
		reportError(err)
	}
}
