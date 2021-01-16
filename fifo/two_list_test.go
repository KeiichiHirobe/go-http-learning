package fifo

import "testing"

func TestTwoListCopy(t *testing.T) {
	l := TwoList{}
	_, l = l.PushBack(1)
	_, l = l.PushBack(2)
	// l2 never change
	l2 := l
	checkTwoList(t, &l, []interface{}{1, 2})
	checkTwoList(t, &l2, []interface{}{1, 2})
	_, l = l.PushBack(3)
	checkTwoList(t, &l, []interface{}{1, 2, 3})
	checkTwoList(t, &l2, []interface{}{1, 2})
	_, l = l.PopFront()
	_, l = l.PopFront()
	checkTwoList(t, &l, []interface{}{3})
	checkTwoList(t, &l2, []interface{}{1, 2})
	_, l = l.PopFront()
	checkTwoList(t, &l, nil)
	checkTwoList(t, &l2, []interface{}{1, 2})
	_, l = l.PushBack(4)
	checkTwoList(t, &l, []interface{}{4})
	checkTwoList(t, &l2, []interface{}{1, 2})
}
