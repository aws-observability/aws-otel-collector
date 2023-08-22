package errors

// ErrorChain is type cast to ChainError for backwards compatibility
// This will be deprecated in future major release
type ErrorChain = ChainError

// ChainError is a linked list of error pointers that point to a parent above and a child below.
type ChainError struct {
	// Head of the linked list
	head error
	// Next node in the linked list
	next error
	// tail node in the linked list
	tail error
}

// Tail is the end of the linked list
func (e *ChainError) Tail() error {
	return e.tail
}

// Head is the start of the linked list
func (e *ChainError) Head() error {
	return e.head
}

// Next is the next node in the linked list
func (e *ChainError) Next() error {
	return e.next
}

// Error returns the error string of the tail of the linked list
func (e *ChainError) Error() string {
	return Cause(e).Error()
}

type errLinkedList interface {
	Head() error
	Next() error
	Tail() error
}

var _ errLinkedList = &ChainError{}
