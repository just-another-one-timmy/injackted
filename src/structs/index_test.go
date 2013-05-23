/*
  Tests for index data structure.
*/

package structs

import (
	"testing"
)

func TestSmokeTest(t *testing.T) {
	const keyword1, keyword2, keyword3 = "ABBA", "Beatles", "Deep Purple"
	const docId1, docId2, docId3 = 1, 123, 321
	// docId1 contains keyword1, keyword2
	// docId2 contains keyword2, keyword3
	// docId3 contains keyword3, keyword1
	var index = NewIndex()
	index.ConnectKeywordDoc(keyword1, docId1)
	index.ConnectKeywordDoc(keyword2, docId1)
	index.ConnectKeywordDoc(keyword2, docId2)
	index.ConnectKeywordDoc(keyword3, docId2)
	index.ConnectKeywordDoc(keyword3, docId3)
	index.ConnectKeywordDoc(keyword1, docId3)
	if !index.IsConnected(keyword1, docId1) {
		t.Logf("%v and %v should be connected.", keyword1, docId1)
		t.Fail()
	}
	if !index.IsConnected(keyword1, docId3) {
		t.Logf("%v and %v should be connected.", keyword1, docId3)
		t.Fail()
	}
	if !index.IsConnected(keyword2, docId1) {
		t.Logf("%v and %v should be connected.", keyword2, docId1)
		t.Fail()
	}
	if !index.IsConnected(keyword2, docId2) {
		t.Logf("%v and %v should be connected.", keyword2, docId2)
		t.Fail()
	}
	if !index.IsConnected(keyword3, docId2) {
		t.Logf("%v and %v should be connected.", keyword3, docId2)
		t.Fail()
	}
	if !index.IsConnected(keyword3, docId3) {
		t.Logf("%v and %v should be connected.", keyword3, docId3)
		t.Fail()
	}
	if index.IsConnected(keyword3, docId1) {
		t.Logf("%v and %v should not be connected.", keyword3, docId1)
		t.Fail()
	}
	if index.IsConnected(keyword1, docId2) {
		t.Logf("%v and %v should not be connected.", keyword1, docId2)
		t.Fail()
	}
	if index.IsConnected(keyword2, docId3) {
		t.Logf("%v and %v should not be connected.", keyword2, docId3)
		t.Fail()
	}
}

func iteratorToSet(iteratorChannel chan interface{}) *Set {
	resultSet := NewSet()
	for item := range iteratorChannel {
		resultSet.Add(item)
	}
	return resultSet
}

func TestIteratorDocsByKeyword(t *testing.T) {
	index := NewIndex()
	const keyword1, keyword2 = "ABBA", "Beatles"
	const docId1, docId2 = 1, 2
	index.ConnectKeywordDoc(keyword1, docId1)

	set := NewSet()
	set.Add(docId1)

	if !iteratorToSet(index.IteratorDocsByKeyword(keyword1)).Equals(set) {
		t.Log("IteratorDocsByKeyword() should iterate over expected set of documents.")
		t.Fail()
	}
	if !iteratorToSet(index.IteratorDocsByKeyword(keyword2)).IsEmpty() {
		t.Log("IteratorDocsByKeyword() should iterate over an empty set for an unused keyword.")
		t.Fail()
	}
}

func TestIteratorKeywordsByDoc(t *testing.T) {
	index := NewIndex()
	const keyword1, keyword2 = "ABBA", "Beatles"
	const docId1, docId2 = 1, 2
	index.ConnectKeywordDoc(keyword1, docId1)

	set := NewSet()
	set.Add(keyword1)

	if !iteratorToSet(index.IteratorKeywordsByDoc(docId1)).Equals(set) {
		t.Log("IteratorKeywordsByDoc() should iterate over expected set of keywords.")
		t.Fail()
	}
	if !iteratorToSet(index.IteratorKeywordsByDoc(docId2)).IsEmpty() {
		t.Log("IteratorKeywordsByDoc() should iterate over an empty set for an unused document.")
		t.Fail()
	}
}

func TestIteratorKeywords(t *testing.T) {
	index := NewIndex()

	if !iteratorToSet(index.IteratorKeywords()).IsEmpty() {
		t.Log("IteratorKeywords() should iterate over an empty set if there are no keywords in the index.")
		t.Fail()
	}

	index.ConnectKeywordDoc("ABBA", "bands")
	index.ConnectKeywordDoc("Beatles", "bands")
	index.ConnectKeywordDoc("Deep Purple", "bands")
	index.ConnectKeywordDoc("Cat", "animals")
	index.ConnectKeywordDoc("Dog", "animals")
	index.ConnectKeywordDoc("Cat", "bands")
	index.ConnectKeywordDoc("Black", "bands")
	index.ConnectKeywordDoc("Black", "colors")
	index.ConnectKeywordDoc("Yellow", "colors")

	expectedSet := NewSet()
	expectedSet.Add("ABBA")
	expectedSet.Add("Beatles")
	expectedSet.Add("Deep Purple")
	expectedSet.Add("Cat")
	expectedSet.Add("Dog")
	expectedSet.Add("Black")
	expectedSet.Add("Yellow")

	if !iteratorToSet(index.IteratorKeywords()).Equals(expectedSet) {
		t.Log("IteratorKeywords() should iterate over expected set of keywords.")
		t.Fail()
	}
}

func TestIteratorDocs(t *testing.T) {
	index := NewIndex()

	if !iteratorToSet(index.IteratorDocs()).IsEmpty() {
		t.Log("IteratorDocs() should iterate over an empty set if there are no docs in the index.")
		t.Fail()
	}

	index.ConnectKeywordDoc("ABBA", "bands")
	index.ConnectKeywordDoc("Beatles", "bands")
	index.ConnectKeywordDoc("Deep Purple", "bands")
	index.ConnectKeywordDoc("Cat", "animals")
	index.ConnectKeywordDoc("Dog", "animals")
	index.ConnectKeywordDoc("Cat", "bands")
	index.ConnectKeywordDoc("Black", "bands")
	index.ConnectKeywordDoc("Black", "colors")
	index.ConnectKeywordDoc("Yellow", "colors")

	expectedSet := NewSet()
	expectedSet.Add("bands")
	expectedSet.Add("animals")
	expectedSet.Add("colors")

	if !iteratorToSet(index.IteratorDocs()).Equals(expectedSet) {
		t.Log("IteratorKeywords() should iterate over expected set of docs.")
		t.Fail()
	}
}
