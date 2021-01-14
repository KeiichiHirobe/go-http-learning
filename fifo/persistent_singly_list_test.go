package fifo

import "testing"

func TestPListCopy(t *testing.T) {
	l := PList{}
	_, l = l.PushBack(1)
	_, l = l.PushBack(2)
	// l2 never change
	l2 := l
	checkPList(t, &l, []interface{}{1, 2})
	checkPList(t, &l2, []interface{}{1, 2})
	_, l = l.PushBack(3)
	checkPList(t, &l, []interface{}{1, 2, 3})
	checkPList(t, &l2, []interface{}{1, 2})
	_, l = l.PopFront()
	_, l = l.PopFront()
	checkPList(t, &l, []interface{}{3})
	checkPList(t, &l2, []interface{}{1, 2})
	_, l = l.PopFront()
	checkPList(t, &l, nil)
	checkPList(t, &l2, []interface{}{1, 2})
	_, l = l.PushBack(4)
	checkPList(t, &l, []interface{}{4})
	checkPList(t, &l2, []interface{}{1, 2})
}
