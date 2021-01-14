package fifo

import "testing"

// copy reference
func TestSListCopyRef(t *testing.T) {
	l := &SList{}
	l.PushBack(1)
	l.PushBack(2)
	l2 := l
	checkList(t, l, []interface{}{1, 2})
	checkList(t, l2, []interface{}{1, 2})
	l.PushBack(3)
	checkList(t, l, []interface{}{1, 2, 3})
	checkList(t, l2, []interface{}{1, 2, 3})
	l.PopFront()
	l.PopFront()
	checkList(t, l, []interface{}{3})
	checkList(t, l2, []interface{}{3})
	l.PopFront()
	checkList(t, l, nil)
	checkList(t, l2, nil)
	l.PushBack(4)
	checkList(t, l, []interface{}{4})
	checkList(t, l2, []interface{}{4})
}

// copy value
// deactivating next link of Element at PopFront in my implementation. Otherwise, we would get result commentouted.
func TestSListCopyValue(t *testing.T) {
	l := SList{}
	l.PushBack(1)
	l.PushBack(2)
	l2 := l
	checkList(t, &l, []interface{}{1, 2})
	checkList(t, &l2, []interface{}{1, 2})
	l.PushBack(3)
	checkList(t, &l, []interface{}{1, 2, 3})
	checkList(t, &l2, []interface{}{1, 2, 3})
	l.PopFront()
	l.PopFront()
	checkList(t, &l, []interface{}{3})
	checkList(t, &l2, []interface{}{1})
	// checkList(t, &l2, []interface{}{1, 2, 3})
	l.PopFront()
	checkList(t, &l, nil)
	checkList(t, &l2, []interface{}{1})
	// checkList(t, &l2, []interface{}{1, 2, 3})
	l.PushBack(4)
	checkList(t, &l, []interface{}{4})
	checkList(t, &l2, []interface{}{1})
	// checkList(t, &l2, []interface{}{1, 2, 3}) not 1,2,3,4
}
