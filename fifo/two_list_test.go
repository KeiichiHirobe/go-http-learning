package fifo

import "testing"

func TestTwoListCopy(t *testing.T) {
	l := TwoList{}
	l.PushBack(1)
	l.PushBack(2)
	// l2 never change
	l2 := l
	checkTwoList(t, &l, []interface{}{1, 2})
	checkTwoList(t, &l2, []interface{}{1, 2})
	l.PushBack(3)
	checkTwoList(t, &l, []interface{}{1, 2, 3})
	checkTwoList(t, &l2, []interface{}{1, 2})
	l.PopFront()
	l.PopFront()
	checkTwoList(t, &l, []interface{}{3})
	checkTwoList(t, &l2, []interface{}{1, 2})
	l.PopFront()
	checkTwoList(t, &l, nil)
	checkTwoList(t, &l2, []interface{}{1, 2})
	l.PushBack(4)
	checkTwoList(t, &l, []interface{}{4})
	checkTwoList(t, &l2, []interface{}{1, 2})
}
