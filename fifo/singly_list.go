package fifo

// implemented refering to https://github.com/golang/go/blob/master/src/container/list/list.go

// SList represents a singley linked list.
// The zero value for List is an empty list ready to use.
type SList struct {
	head *Element
	tail *Element
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *SList) PushBack(v interface{}) *Element {
	e := &Element{
		Value: v,
	}
	if l.head == nil {
		l.tail = e
		l.head = e
	} else {
		l.tail.next = e
		l.tail = e
	}
	return e
}

// PopFront remove removes e from head of list l and returns e.
// Return nil if empty.
func (l *SList) PopFront() *Element {
	head := l.head

	if head == nil {
		// zero length
		return nil
	} else if head == l.tail {
		// one length
		l.head = nil
		l.tail = nil
		// avoid memory leaks
		head.next = nil
	} else {
		l.head = head.next
		// avoid memory leaks
		head.next = nil
	}
	return head
}
