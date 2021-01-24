package fifo

import (
	"sync"
)

func nop() {}

var (
	testHookEvaluated = nop
)

// BankerList represents a banker queue.
// BankerList is persistent list.
// The zero value for List is an empty list ready to use.
// see also http://www.kmonos.net/pub/Presen/PFDS.pdf
type BankerList struct {
	// head is a list of *headL.
	// head is a persistent queue.
	head TwoList
	// reversed is a list after evaluation of reverse.
	// at PopFront, first see reversed, and if empty PopFront from head and reverse the list and set to reversed.
	reversed *Element
	// the number of element like tailN, not the number of element for head.
	headN int
	tail  *Element
	tailN int
}

// headL represents a singly linked list.
// reverse lazily and result is memorized.
type headL struct {
	head *Element
	once sync.Once
}

// eval may be called concurrently.
func (l *headL) eval() {
	l.once.Do(
		func() {
			// reverse list
			var prev *Element
			for e := l.head; e != nil; e = e.next {
				el := &Element{
					Value: e.Value,
				}
				if prev != nil {
					el.next = prev
				}
				prev = el
			}
			l.head = prev
			// for test.
			testHookEvaluated()
		},
	)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *BankerList) PushBack(v interface{}) *Element {
	e := &Element{
		Value: v,
	}
	if l.tail == nil {
		l.tail = e
	} else {
		e.next = l.tail
		l.tail = e
	}
	l.tailN++
	l.Chk()
	return e
}

// PopFront removes e from head of list l and returns e.
// Return nil if empty.
func (l *BankerList) PopFront() *Element {
	if l.reversed == nil {
		el := l.head.PopFront()
		if el == nil {
			return nil
		}
		he := el.Value.(*headL)
		he.eval()
		l.reversed = he.head
	}
	if l.reversed == nil {
		panic("Must not be empty")
	}
	el := l.reversed
	l.reversed = l.reversed.next
	l.headN--
	l.Chk()
	return el
}

// Check and move from tail to head if need.
func (l *BankerList) Chk() {
	if l.tailN == l.headN+1 {
		l.headN += l.tailN
		l.tailN = 0
		// reverse lazily
		lazyL := &headL{
			head: l.tail,
		}
		l.tail = nil
		l.head.PushBack(lazyL)
	}
}
