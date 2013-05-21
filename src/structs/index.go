package structs

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

func (index *Index) IsConnected(keyword interface{}, doc interface{}) bool {
	if _, present := index.docsByKeyword[keyword]; !present {
		return false
	}
	return index.docsByKeyword[keyword].ItemPresent(doc)
}
