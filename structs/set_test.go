/*
 Tests for set data structure.
*/

package structs

import (
	"testing"
)

func TestEmptySet(t *testing.T) {
	var emptySet = NewSet()
	if emptySet.Size() > 0 {
		t.Log("Size() should return 0 for empty set.")
		t.Fail()
	}
	if !emptySet.IsEmpty() {
		t.Log("IsEmpty() should return true for empty set.")
		t.Fail()
	}
}

func testIsEmpty(t *testing.T) {
	var set = NewSet()
	set.Add(1)
	if set.IsEmpty() {
		t.Log("IsEmpty() should return false for non-empty set.")
		t.Fail()
	}
}

func TestAddOneItem(t *testing.T) {
	var set = NewSet()
	const item1, item2 = 1, 2

	set.Add(item1)
	if set.Size() != 1 {
		t.Log("Size() should return 1 for 1-element set.")
		t.Fail()
	}
	if !set.ItemPresent(item1) {
		t.Log("ItemPresent() should return true for present item.")
		t.Fail()
	}

	set.Add(item2)
	if set.Size() != 2 {
		t.Log("Size() should return 2 for 2-element set.")
		t.Fail()
	}
	if !set.ItemPresent(item2) {
		t.Log("ItemPresent() should return true for present item.")
		t.Fail()
	}

	set.Add(item1)
	if set.Size() != 2 {
		t.Log("Adding element that's present should not change the size.")
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	var set = NewSet()
	const item1, item2 = 1, 2
	set.Add(item1)
	set.Remove(item1)
	if set.ItemPresent(item1) {
		t.Log("Item should have been deleted")
		t.Fail()
	}
	if set.Size() != 0 {
		t.Log("Size() returns wrong value after deletion of element that was present.")
		t.Fail()
	}
	set.Remove(item1)
	if set.Size() != 0 {
		t.Log("Size() returns wrong value after deletion of element that was not present.")
		t.Fail()
	}
}

func TestEquals(t *testing.T) {
	var set1, set2 = NewSet(), NewSet()
	const item1, item2 = 1, 2
	if !set1.Equals(set2) {
		t.Log("Equals() should return true for equal sets.")
		t.Fail()
	}
	// In this function we also Check that Equals() is commutative.
	if !set2.Equals(set1) {
		t.Log("Equals() should be commutative")
		t.Fail()
	}
	// And reflexive.
	if !set1.Equals(set1) {
		t.Log("Equals() should be reflexive.")
		t.Fail()
	}

	set1.Add(item1)
	set2.Add(item2)
	if set1.Equals(set2) {
		t.Log("Equals() should return false for non-equal sets.")
		t.Fail()
	}
	if set2.Equals(set1) {
		t.Log("Equals() should be commutative")
		t.Fail()
	}
	if !set1.Equals(set1) {
		t.Log("Equals() should be reflexive.")
		t.Fail()
	}

	set1.Add(item2)
	set2.Add(item1)
	if !set1.Equals(set2) {
		t.Log("Equals() should return true for equal sets.")
		t.Fail()
	}
	if !set2.Equals(set1) {
		t.Log("Equals() should be commutative")
		t.Fail()
	}
	if !set1.Equals(set1) {
		t.Log("Equals() should be reflexive.")
		t.Fail()
	}
}

func TestIntersection(t *testing.T) {
	//TODO(iaroslav): write test.
}

func TestUnion(t *testing.T) {
	//TODO(iaroslav): write test.
}

func TestDifference(t *testing.T) {
	//TODO(iaroslav): write test.
}
