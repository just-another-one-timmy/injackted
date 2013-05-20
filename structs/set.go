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

func NewSet() *Set {
	return &Set{make(map[interface{}]struct{}), 0}
}

func (set *Set) Size() int {
	return set.size
}

func (set *Set) IsEmpty() bool {
	return set.Size() == 0
}

// Add item to the set.
func (set *Set) Add(item interface{}) {
	if !set.ItemPresent(item) {
		set.internal[item] = struct{}{}
		set.size += 1
	}
}

// Check if item is present in the set.
func (set *Set) ItemPresent(item interface{}) bool {
	_, present := set.internal[item]
	return present
}

// Remove item from the set.
// Works even when item is not present.
func (set *Set) Remove(item interface{}) {
	if set.ItemPresent(item) {
		delete(set.internal, item)
		set.size -= 1
	}
}
