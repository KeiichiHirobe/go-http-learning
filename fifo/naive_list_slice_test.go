package fifo

import "testing"

func TestNaiveListSliceCopy(t *testing.T) {
	l := NaiveListSlice{}
	l.PushBack(1)
	l.PushBack(2)
	l2 := l
	checkNaiveListSlice(t, &l, []interface{}{1, 2})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PopFront()
	checkEqualEl(t, l.Next(), interface{}(2))
	checkNaiveListSlice(t, &l, []interface{}{2})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PopFront()
	checkNaiveListSlice(t, &l, nil)
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	// when empty, return nil
	l.PopFront()
	checkEqualEl(t, l.Next(), nil)
	checkNaiveListSlice(t, &l, nil)
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PushBack(3)
	l.PopFront()
	checkNaiveListSlice(t, &l, []interface{}{})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	checkNaiveListSlice(t, &l, []interface{}{4, 5, 6})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
}

func TestNaiveListSliceCopyOverWrite(t *testing.T) {
	l := NaiveListSlice{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l2 := l
	checkNaiveListSlice(t, &l, []interface{}{1, 2, 3})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2, 3})
	l.PushBack(4)
	checkNaiveListSlice(t, &l, []interface{}{1, 2, 3, 4})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2, 3})
	l2.PushBack(5)
	// PushBack to l2 overwrite l
	checkNaiveListSlice(t, &l, []interface{}{1, 2, 3, 5})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2, 3, 5})
}
