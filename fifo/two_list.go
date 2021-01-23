package fifo

// http://www.kmonos.net/pub/Presen/PFDS.pdf

// TwoList represents a two list queue.
// TwoList is persistent list.
// The zero value for List is an empty list ready to use.
type TwoList struct {
	head *Element
	tail *Element
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *TwoList) PushBack(v interface{}) *Element {
	e := &Element{
		Value: v,
	}
	if l.tail == nil {
		l.tail = e
	} else {
		e.next = l.tail
		l.tail = e
	}
	return e
}

// PopFront removes e from head of list l and returns e.
// Return nil if empty.
func (l *TwoList) PopFront() *Element {
	return l.seekFront(true)
}

// PeekFront peek head of list l and returns e. Never remove.
// Reverse may occur if need.
func (l *TwoList) PeekFront() *Element {
	return l.seekFront(false)
}

// seekFront seek head of list l and returns e.
// If remove is true, remove e from list.
func (l *TwoList) seekFront(remove bool) *Element {
	head := l.head
	if head != nil {
		if remove {
			l.head = head.next
		}
		return head
	}

	if l.tail == nil {
		return nil
	}

	// list which starts from head is empty, so
	// reverse list which starts from tail and attach to head(not tail)
	var prev *Element
	for e := l.tail; e != nil; e = e.next {
		el := &Element{
			Value: e.Value,
		}
		if prev != nil {
			el.next = prev
		}
		prev = el
	}
	// reset list
	l.tail = nil
	if !remove {
		l.head = prev
	} else {
		l.head = prev.next
	}
	return prev
}
