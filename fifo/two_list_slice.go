package fifo

// TwoListSlice are based on wantConnQueue implemented by Russ Cox
// https://github.com/golang/go/blob/682a1d2176b02337460aeede0ff9e49429525195/src/net/http/transport.go#L1242-L1306

// TwoListSlice represents a two list queue without the overhead of reversing the list when swapping stages.
// TwoListSlice is NOT a persistent queue.
// The zero value for List is an empty list ready to use.
type TwoListSlice struct {
	head    []interface{}
	headPos int
	tail    []interface{}
}

// Len returns the number of items in the queue.
func (l *TwoListSlice) Len() int {
	return len(l.head) - l.headPos + len(l.tail)
}

// PushBack adds v to the back of the queue.
func (l *TwoListSlice) PushBack(v interface{}) {
	l.tail = append(l.tail, v)
}

// PopFront removes and returns the v at the front of the queue.
func (l *TwoListSlice) PopFront() interface{} {
	if l.headPos >= len(l.head) {
		if len(l.tail) == 0 {
			return nil
		}
		// Pick up tail as new head, clear tail.
		l.head, l.headPos, l.tail = l.tail, 0, l.head[:0]
	}
	v := l.head[l.headPos]
	// avoid memory leaks
	l.head[l.headPos] = nil
	l.headPos++
	return v
}
