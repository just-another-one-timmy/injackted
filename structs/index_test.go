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
	index.connectKeywordDoc(keyword1, docId1)
	index.connectKeywordDoc(keyword2, docId1)
	index.connectKeywordDoc(keyword2, docId2)
	index.connectKeywordDoc(keyword3, docId2)
	index.connectKeywordDoc(keyword3, docId3)
	index.connectKeywordDoc(keyword1, docId3)
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
