package db

type Iterator interface {
	// Returns true iff the iterator is positioned at a valid node.
	Valid() bool

	// Returns the key at the current position.
	// REQUIRES: Valid()
	Key() interface{}

	// Advances to the next position.
	// REQUIRES: Valid()
	Next()

	// Advances to the previous position.
	// REQUIRES: Valid()
	Prev()

	// Advance to the first entry with a key >= target
	Seek(target []byte)

	// Position at the first entry in list.
	// Final state of iterator is Valid() iff list is not empty.
	SeekToFirst()

	// Position at the last entry in list.
	// Final state of iterator is Valid() iff list is not empty.
	SeekToLast()
}
