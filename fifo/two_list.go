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

// PopFront remove removes e from head of list l and returns e.
// Return nil if empty.
func (l *TwoList) PopFront() *Element {
	head := l.head
	if head != nil {
		l.head = head.next
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
	l.head = prev.next
	return prev
}
