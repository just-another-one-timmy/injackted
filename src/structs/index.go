package structs

import (
	"io"
	"strconv"
)

type Index struct {
	docsByKeyword, keywordsByDoc map[interface{}]*Set
}

func NewIndex() *Index {
	return &Index{make(map[interface{}]*Set),
		make(map[interface{}]*Set)}
}

// Connects keyword and doc.
// Neither of them should be present prior to connecting.
func (index *Index) ConnectKeywordDoc(keyword interface{}, doc interface{}) {
	addItem(index.docsByKeyword, keyword, doc)
	addItem(index.keywordsByDoc, doc, keyword)
}

// Adds value to the set returned from map by key.
// If key is not present, it will be added and empty set will be created.
func addItem(setsByKey map[interface{}]*Set, key interface{}, value interface{}) {
	if _, present := setsByKey[key]; !present {
		setsByKey[key] = NewSet()
	}
	setsByKey[key].Add(value)
}

// Checks if given keyword and document are connected in the index.
func (index *Index) IsConnected(keyword interface{}, doc interface{}) bool {
	if _, present := index.docsByKeyword[keyword]; !present {
		return false
	}
	return index.docsByKeyword[keyword].ItemPresent(doc)
}

// Returns an iterator over documents connected to given keyword in the index.
func (index *Index) IteratorDocsByKeyword(keyword interface{}) chan interface{} {
	if _, present := index.docsByKeyword[keyword]; !present {
		return NewSet().Iterator()
	}
	return index.docsByKeyword[keyword].Iterator()
}

// Returns an iterator over keywords connected to given doc in the index.
func (index *Index) IteratorKeywordsByDoc(doc interface{}) chan interface{} {
	if _, present := index.keywordsByDoc[doc]; !present {
		return NewSet().Iterator()
	}
	return index.keywordsByDoc[doc].Iterator()
}

func iteratorOverMapKeys(mapToIterateOver map[interface{}]*Set) chan interface{} {
	resultChannel := make(chan interface{})
	go func() {
		for key, _ := range mapToIterateOver {
			resultChannel <- key
		}
		close(resultChannel)
	}()
	return resultChannel
}

// Returns an iterator over all keywords present in the index.
func (index *Index) IteratorKeywords() chan interface{} {
	return iteratorOverMapKeys(index.docsByKeyword)
}

// Returns an iterator over all docs present in the index.
func (index *Index) IteratorDocs() chan interface{} {
	return iteratorOverMapKeys(index.keywordsByDoc)
}

func (index *Index) Print(w io.Writer) (n int, err error) {
	resultN := 0
	for keyword := range index.IteratorKeywords() {
		stringKeyword, _ := keyword.(string)
		if n, err := w.Write([]byte(stringKeyword + " ")); err != nil {
			return n + resultN, err
		} else {
			resultN += n
		}
		docsCount := 0
		for _ = range index.IteratorDocsByKeyword(keyword) {
			docsCount++
		}
		w.Write([]byte(strconv.Itoa(docsCount) + "\n"))
		for doc := range index.IteratorDocsByKeyword(keyword) {
			stringDoc, _ := doc.(string)
			if n, err := w.Write([]byte("\t" + stringDoc + "\n")); err != nil {
				return n + resultN, err
			} else {
				resultN += n
			}
		}
	}
	return resultN, nil
}
