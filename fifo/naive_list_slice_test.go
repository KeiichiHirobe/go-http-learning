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
	checkNaiveListSlice(t, &l, []interface{}{2})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PopFront()
	checkNaiveListSlice(t, &l, nil)
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	l.PushBack(3)
	l.PopFront()
	checkNaiveListSlice(t, &l, []interface{}{})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
	// Here, array memory pointed by l2's tail is overwritten by l1's PushBack
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	checkNaiveListSlice(t, &l, []interface{}{4, 5, 6})
	checkNaiveListSlice(t, &l2, []interface{}{1, 2})
}
