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

func TestTwoListCopy2(t *testing.T) {
	l := TwoList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l2 := l
	checkTwoList(t, &l, []interface{}{1, 2, 3})
	checkTwoList(t, &l2, []interface{}{1, 2, 3})
	l.PushBack(4)
	checkTwoList(t, &l, []interface{}{1, 2, 3, 4})
	checkTwoList(t, &l2, []interface{}{1, 2, 3})
	l2.PushBack(5)
	checkTwoList(t, &l, []interface{}{1, 2, 3, 4})
	checkTwoList(t, &l2, []interface{}{1, 2, 3, 5})
	checkEqualEl(t, l.PopFront().Value, 1)
	checkEqualEl(t, l.PopFront().Value, 2)
	checkEqualEl(t, l.PopFront().Value, 3)
	checkEqualEl(t, l.PopFront().Value, 4)
	checkEqualEl(t, l2.PopFront().Value, 1)
	checkEqualEl(t, l2.PopFront().Value, 2)
	checkEqualEl(t, l2.PopFront().Value, 3)
	checkEqualEl(t, l2.PopFront().Value, 5)
}
