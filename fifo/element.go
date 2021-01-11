package fifo

// Element is an element of a singley linked list.
type Element struct {
	next *Element
	// The value stored with this element.
	Value interface{}
}
