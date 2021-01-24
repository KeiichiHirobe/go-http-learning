package fifo

import (
	"bytes"
	"fmt"
	"testing"
)

var (
	setEvaluatedHook = hookSetter(&testHookEvaluated)
)

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

func Test2BankerList(t *testing.T) {
	l := BankerList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l2 := l
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	// r[1]r[3,2]r[7,6,5,4]
	l.PushBack(7)

	l2.PushBack(14)
	l2.PushBack(15)
	l2.PushBack(16)
	// r[1]r[3,2]r[17,16,15,14]
	l2.PushBack(17)

	checkEqualEl(t, l.PopFront().Value, 1)
	// 3 r[7,6,5,4]
	checkEqualEl(t, l.PopFront().Value, 2)
	checkEqualEl(t, l2.PopFront().Value, 1)
	checkEqualEl(t, l2.PopFront().Value, 2)
	// 3 r[17,16,15,14]
	checkEqualEl(t, l2.PopFront().Value, 3)
	checkEqualEl(t, l2.PopFront().Value, 14)

	checkEqualEl(t, l.PopFront().Value, 3)
	// fail test if we update next link of element which has value 3 to r[17,16,15,14] because link to r[7,6,5,4] is lost.
	// FIXME want not 14 but 4
	checkEqualEl(t, l.PopFront().Value, 14)
}

func TestBankerListEval(t *testing.T) {
	var (
		logbuf bytes.Buffer
	)

	logf := func(format string, args ...interface{}) {
		_, _ = fmt.Fprintf(&logbuf, format, args...)
	}
	setEvaluatedHook(func() {
		logf("Evaluated")
	})
	defer setEvaluatedHook(nil)

	l := BankerList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)
	l.PushBack(8)
	l.PushBack(9)
	l.PushBack(10)
	l.PushBack(11)
	l.PushBack(12)
	l.PushBack(13)
	l.PushBack(14)
	l.PushBack(15)
	l2 := l
	// never evaluated until here
	checkEqualLogBuf(t, logbuf.String(), "")
	l.PopFront()
	// [1] -> [1]
	checkEqualLogBuf(t, logbuf.String(), "Evaluated")
	logbuf.Reset()
	l.PopFront()
	// [3,2] -> [2,3]
	checkEqualLogBuf(t, logbuf.String(), "Evaluated")
	logbuf.Reset()
	l.PopFront()
	checkEqualLogBuf(t, logbuf.String(), "")
	e := l.PopFront()
	checkEqualEl(t, e.Value, 4)
	// [7,6,5,4] -> [4,5,6,7]
	checkEqualLogBuf(t, logbuf.String(), "Evaluated")
	logbuf.Reset()
	l.PopFront()
	l.PopFront()
	l.PopFront()
	checkEqualLogBuf(t, logbuf.String(), "")
	e = l.PopFront()
	checkEqualEl(t, e.Value, 8)
	// [15,14,..,8] -> [8,9,..,15]
	checkEqualLogBuf(t, logbuf.String(), "Evaluated")
	logbuf.Reset()
	l.PopFront()
	l.PopFront()
	l.PopFront()
	l.PopFront()
	l.PopFront()
	l.PopFront()
	e = l.PopFront()
	checkEqualEl(t, e.Value, 15)
	checkEqualLogBuf(t, logbuf.String(), "")

	for i := 0; i < 15; i++ {
		e = l2.PopFront()
	}
	checkEqualEl(t, e.Value, 15)
	checkEqualLogBuf(t, logbuf.String(), "")
}
