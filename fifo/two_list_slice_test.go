package fifo

import "testing"

func TestTwoListSliceCopy(t *testing.T) {
	l := TwoListSlice{}
	l.PushBack(1)
	l.PushBack(2)
	l2 := l
	checkTwoListSlice(t, &l, []interface{}{1, 2})
	checkTwoListSlice(t, &l2, []interface{}{1, 2})
	l.PushBack(3)
	checkTwoListSlice(t, &l, []interface{}{1, 2, 3})
	checkTwoListSlice(t, &l2, []interface{}{1, 2})
	l.PopFront()
	l.PopFront()
	checkTwoListSlice(t, &l, []interface{}{3})
	checkTwoListSlice(t, &l2, []interface{}{1, 2})
	l.PopFront()
	checkTwoListSlice(t, &l, nil)
	checkTwoListSlice(t, &l2, []interface{}{1, 2})
	l.PushBack(4)
	checkTwoListSlice(t, &l, []interface{}{4})
	checkTwoListSlice(t, &l2, []interface{}{1, 2})
}
