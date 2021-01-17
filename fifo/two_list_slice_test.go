package fifo

import "testing"

func TestTwoListSliceCopy(t *testing.T) {
	l := TwoListSlice{}
	l.PushBack(1)
	l.PushBack(2)
	l2 := l
	checkTwoListSlice(t, &l, []interface{}{1, 2})
	checkTwoListSlice(t, &l2, []interface{}{1, 2})
	l.PopFront()
	checkTwoListSlice(t, &l, []interface{}{2})
	checkTwoListSlice(t, &l2, []interface{}{nil, 2})
	l.PopFront()
	checkTwoListSlice(t, &l, nil)
	checkTwoListSlice(t, &l2, []interface{}{nil, nil})
	l.PushBack(3)
	l.PopFront()
	checkTwoListSlice(t, &l, []interface{}{})
	checkTwoListSlice(t, &l2, []interface{}{nil, nil})
	// Here, array memory pointed by l2's tail is overwritten by l1's PushBack
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	checkTwoListSlice(t, &l, []interface{}{4, 5, 6})
	checkTwoListSlice(t, &l2, []interface{}{4, 5})
}
