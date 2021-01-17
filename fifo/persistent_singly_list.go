package fifo

// PList represents a persistent singley linked list.
// The zero value for List is an empty list ready to use.
type PList struct {
	head *Element
	tail *Element
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *PList) PushBack(v interface{}) *Element {
	l.Clone()
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
func (l *PList) PopFront() *Element {
	l.Clone()
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

func (l *PList) Clone() {
	var head, prev *Element
	for e := l.head; e != nil; e = e.next {
		el := &Element{
			Value: e.Value,
		}
		if prev == nil {
			head = el
		} else {
			prev.next = el
		}
		prev = el
	}
	l.head = head
	l.tail = prev
}
