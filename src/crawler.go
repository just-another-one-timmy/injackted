B1;3402;0cpackage main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"os"
	"structs"
)

type DocKeywordPair struct {
	Doc     string
	Keyword string
}

func reportError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err.Error())
}

func processFile(fileName string, c chan DocKeywordPair, w chan string) {
	// Inform receiver that current goroutine is done.
	defer func() {
		w <- fileName
	}()

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
	w := make(chan string)
	goroutinesCount := len(fileNames)
	for _, fileName := range fileNames {
		go processFile(fileName, c, w)
	}
	for goroutinesCount > 0 {
		select {
		case docKeywordPair := <-c:
			index.ConnectKeywordDoc(docKeywordPair.Keyword,
				docKeywordPair.Doc)
		case name := <-w:
			percentage := (1.0 - float64(goroutinesCount)/float64(len(fileNames))) * 100
			fmt.Printf("Processed file %v\t%6.2f%%\n", name, percentage)
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
func saveIndex(index *structs.Index, w io.Writer) error {
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(structs.IndexToGobIndex(index))
	return err
}

func main() {
	var outputFileName string
	flag.StringVar(&outputFileName, "o", "", "File name to write index to.")
	flag.Parse()

	if len(outputFileName) == 0 {
		fmt.Fprintf(os.Stderr, "Output file not specified. Use -o option.\n")
		return
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		reportError(err)
		return
	}

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

	if err = saveIndex(index, outputFile); err != nil {
		reportError(err)
	}
}
