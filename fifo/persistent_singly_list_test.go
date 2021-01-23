package fifo

import "testing"

func TestPListCopy(t *testing.T) {
	l := PList{}
	l.PushBack(1)
	l.PushBack(2)
	l2 := l
	checkPList(t, &l, []interface{}{1, 2})
	checkPList(t, &l2, []interface{}{1, 2})
	l.PushBack(3)
	checkPList(t, &l, []interface{}{1, 2, 3})
	checkPList(t, &l2, []interface{}{1, 2})
	l.PopFront()
	l.PopFront()
	checkPList(t, &l, []interface{}{3})
	checkPList(t, &l2, []interface{}{1, 2})
	l.PopFront()
	checkPList(t, &l, nil)
	checkPList(t, &l2, []interface{}{1, 2})
	l.PushBack(4)
	checkPList(t, &l, []interface{}{4})
	checkPList(t, &l2, []interface{}{1, 2})
}

func TestPlistCopy2(t *testing.T) {
	l := PList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l2 := l
	checkPList(t, &l, []interface{}{1, 2, 3})
	checkPList(t, &l2, []interface{}{1, 2, 3})
	l.PushBack(4)
	checkPList(t, &l, []interface{}{1, 2, 3, 4})
	checkPList(t, &l2, []interface{}{1, 2, 3})
	l2.PushBack(5)
	checkPList(t, &l, []interface{}{1, 2, 3, 4})
	checkPList(t, &l2, []interface{}{1, 2, 3, 5})
}
