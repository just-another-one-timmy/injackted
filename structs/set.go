/*
 Package provides basic data structures to be used in the client program.
*/

package structs

/*
 Set data structure.
*/
type Set struct {
	internal map[interface{}]struct{}
	size     int
}

// Creates new empty set.
func NewSet() *Set {
	return &Set{make(map[interface{}]struct{}), 0}
}

// Returns size of the set.
func (set *Set) Size() int {
	return set.size
}

// Returns true if the set is empty.
func (set *Set) IsEmpty() bool {
	return set.Size() == 0
}

// Adds item to the set.
func (set *Set) Add(item interface{}) {
	if !set.ItemPresent(item) {
		set.internal[item] = struct{}{}
		set.size += 1
	}
}

// Checks if item is present in the set.
func (set *Set) ItemPresent(item interface{}) bool {
	_, present := set.internal[item]
	return present
}

// Removes item from the set.
// Works even when item is not present.
func (set *Set) Remove(item interface{}) {
	if set.ItemPresent(item) {
		delete(set.internal, item)
		set.size -= 1
	}
}
