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
	var set1, set2, set3, emptySet = NewSet(), NewSet(), NewSet(), NewSet()
	const item1, item2, item3 = "ABBA", "Beatles", "Deep Purple"
	set1.Add(item1)
	set1.Add(item2)
	set2.Add(item2)
	set2.Add(item3)
	set3.Add(item2)
	// Intersection should be commutative.
	if !set1.Intersection(set2).Equals(set2.Intersection(set1)) {
		t.Log("Intersection() should be commutative.")
		t.Fail()
	}
	// Intersection should be reflexive.
	if !set1.Intersection(set1).Equals(set1) {
		t.Log("Intersection() should be reflexive.")
	}
	if !set1.Intersection(set2).Equals(set3) {
		t.Log("Intersection() should return expected set.")
		t.Fail()
	}
	if !set3.Intersection(emptySet).IsEmpty() {
		t.Log("Intersection() with empty set should be empty.")
		t.Fail()
	}
}

func TestUnion(t *testing.T) {
	var set1, set2, set3, emptySet = NewSet(), NewSet(), NewSet(), NewSet()
	const item1, item2, item3 = "ABBA", "Beatles", "Deep Purple"
	set1.Add(item1)
	set1.Add(item2)
	set2.Add(item2)
	set2.Add(item3)
	set3.Add(item1)
	set3.Add(item2)
	set3.Add(item3)
	// Union should be commutative.
	if !set1.Union(set2).Equals(set2.Union(set1)) {
		t.Log("Union() should be commutative.")
		t.Fail()
	}
	// Union should be reflexive.
	if !set1.Union(set1).Equals(set1) {
		t.Log("Union() should be reflexive.")
	}
	if !set1.Union(set2).Equals(set3) {
		t.Log("Union() should return expected set.")
		t.Fail()
	}
	if !set3.Union(emptySet).Equals(set3) {
		t.Log("Union() with empty set should return original set.")
		t.Fail()
	}
}

func TestDifference(t *testing.T) {
	var set1, set2, set3, emptySet = NewSet(), NewSet(), NewSet(), NewSet()
	const item1, item2, item3 = "ABBA", "Beatles", "Deep Purple"
	set1.Add(item1)
	set1.Add(item2)
	set2.Add(item2)
	set2.Add(item3)
	set3.Add(item1)
	// Difference between same sets is empty set.
	if !set1.Difference(set1).IsEmpty() {
		t.Log("Difference() with myself should be empty.")
		t.Fail()
	}
	if !set1.Difference(set2).Equals(set3) {
		t.Log("Difference() should return expected set.")
		t.Fail()
	}
	if !set1.Difference(emptySet).Equals(set1) {
		t.Log("Difference() with empty set should return original set.")
		t.Fail()
	}
}

func TestIterator(t *testing.T) {
	var set1, set2 = NewSet(), NewSet()
	const item1, item2 = 123, "Deep Purple"
	set1.Add(item1)
	set1.Add(item2)
	for item := range set1.Iterator() {
		set2.Add(item)
	}
	if !set1.Equals(set2) {
		t.Log("Iterator() should be iterating over all values.")
		t.Fail()
	}
}
