package fifo

import "testing"

func TestNaiveListSliceCopy(t *testing.T) {
	l := NaiveListSlice{}
	l.PushBack(1)
	l.PushBack(2)
	l2 := l
	checkNaiveListSlice(t, &l, []interface{}{1, 2})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PushBack(3)
	checkNaiveListSlice(t, &l, []interface{}{1, 2, 3})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PopFront()
	l.PopFront()
	checkNaiveListSlice(t, &l, []interface{}{3})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PopFront()
	checkNaiveListSlice(t, &l, nil)
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PushBack(4)
	checkNaiveListSlice(t, &l, []interface{}{4})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
}
