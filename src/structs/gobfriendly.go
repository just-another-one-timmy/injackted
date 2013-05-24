package structs

type GobSet struct {
	Internal map[interface{}]struct{}
}

type GobIndex struct {
	DocsByKeyword, KeywordsByDoc map[interface{}]*GobSet
}

func IndexToGobIndex(i *Index) *GobIndex{
	result := &GobIndex{make(map[interface{}]*GobSet),
		make(map[interface{}]*GobSet)}
	for doc, keywords := range i.keywordsByDoc {
		result.KeywordsByDoc[doc] = SetToGobSet(keywords)
	}
	for keyword, docs := range i.docsByKeyword {
		result.DocsByKeyword[keyword] = SetToGobSet(docs)
	}
	return result
}

func GobIndexToIndex(g *GobIndex) *Index{
	result := NewIndex()
	for doc, keywords := range g.KeywordsByDoc {
		for keyword, _ := range keywords.Internal {
			result.ConnectKeywordDoc(keyword, doc)
		}
	}
	return result
}

func SetToGobSet(s *Set) *GobSet{
	result := &GobSet{make(map[interface{}]struct{})}
	for item := range s.Iterator() {
		result.Internal[item] = struct{}{}
	}
	return result
}

func GobSetToSet(g *GobSet) *Set{
	result := NewSet()
	for key, _ := range g.Internal {
		result.Add(key)
	}
	return result
}