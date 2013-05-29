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
func (i *Index) ConnectKeywordDoc(keyword interface{}, doc interface{}) {
	addItem(i.docsByKeyword, keyword, doc)
	addItem(i.keywordsByDoc, doc, keyword)
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
func (i *Index) IsConnected(keyword interface{}, doc interface{}) bool {
	if _, present := i.docsByKeyword[keyword]; !present {
		return false
	}
	return i.docsByKeyword[keyword].ItemPresent(doc)
}

// Returns an iterator over documents connected to given keyword in the index.
func (i *Index) IteratorDocsByKeyword(keyword interface{}) <-chan interface{} {
	if _, present := i.docsByKeyword[keyword]; !present {
		return NewSet().Iterator()
	}
	return i.docsByKeyword[keyword].Iterator()
}

// Returns an iterator over keywords connected to given doc in the index.
func (i *Index) IteratorKeywordsByDoc(doc interface{}) <-chan interface{} {
	if _, present := i.keywordsByDoc[doc]; !present {
		return NewSet().Iterator()
	}
	return i.keywordsByDoc[doc].Iterator()
}

func iteratorOverMapKeys(mapToIterateOver map[interface{}]*Set) <-chan interface{} {
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
func (i *Index) IteratorKeywords() <-chan interface{} {
	return iteratorOverMapKeys(i.docsByKeyword)
}

// Returns an iterator over all docs present in the index.
func (i *Index) IteratorDocs() <-chan interface{} {
	return iteratorOverMapKeys(i.keywordsByDoc)
}

func (i *Index) Print(w io.Writer) (n int, err error) {
	resultN := 0
	for keyword := range i.IteratorKeywords() {
		stringKeyword, _ := keyword.(string)
		if n, err := w.Write([]byte(stringKeyword + " ")); err != nil {
			return n + resultN, err
		} else {
			resultN += n
		}
		docsCount := 0
		for _ = range i.IteratorDocsByKeyword(keyword) {
			docsCount++
		}
		w.Write([]byte(strconv.Itoa(docsCount) + "\n"))
		for doc := range i.IteratorDocsByKeyword(keyword) {
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
