package fifo

import "testing"

func TestBankerList(t *testing.T) {
	l := BankerList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l2 := l
	l.PushBack(4)
	l2.PushBack(5)
	checkEqualEl(t, l.PopFront().Value, 1)
	checkEqualEl(t, l.PopFront().Value, 2)
	checkEqualEl(t, l.PopFront().Value, 3)
	checkEqualEl(t, l.PopFront().Value, 4)
	checkEqualEl(t, l2.PopFront().Value, 1)
	checkEqualEl(t, l2.PopFront().Value, 2)
	checkEqualEl(t, l2.PopFront().Value, 3)
	checkEqualEl(t, l2.PopFront().Value, 5)
}
