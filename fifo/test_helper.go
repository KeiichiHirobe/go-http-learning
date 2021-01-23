package fifo

import "testing"

func checkList(t *testing.T, l *SList, es []interface{}) {
	t.Helper()
	i := 0
	for e := l.head; e != nil; e = e.next {
		le := e.Value.(int)
		if le != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
	if len(es) != i {
		t.Errorf("length = %d, want %d", i, len(es))
	}
}

func checkPList(t *testing.T, l *PList, es []interface{}) {
	t.Helper()
	i := 0
	for e := l.head; e != nil; e = e.next {
		le := e.Value.(int)
		if le != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
	if len(es) != i {
		t.Errorf("length = %d, want %d", i, len(es))
	}
}

func checkTwoList(t *testing.T, l *TwoList, es []interface{}) {
	t.Helper()
	i := 0
	for e := l.head; e != nil; e = e.next {
		le := e.Value.(int)
		if le != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
	var vlist []interface{}
	for e := l.tail; e != nil; e = e.next {
		vlist = append([]interface{}{e.Value}, vlist...)
	}
	for _, v := range vlist {
		if v != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, v, es[i])
		}
		i++
	}
	if len(es) != i {
		t.Errorf("length = %d, want %d", i, len(es))
	}
}

func checkTwoListSlice(t *testing.T, l *TwoListSlice, es []interface{}) {
	t.Helper()
	var vlist []interface{}
	vlist = append(l.head[l.headPos:], l.tail...)
	for i, v := range vlist {
		if v != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, v, es[i])
		}
	}
	if len(es) != l.Len() {
		t.Errorf("length = %d, want %d", l.Len(), len(es))
	}
}

func checkNaiveListSlice(t *testing.T, l *NaiveListSlice, es []interface{}) {
	t.Helper()
	for i, v := range *l {
		if v != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, v, es[i])
		}
	}
	if len(es) != l.Len() {
		t.Errorf("length = %d, want %d", l.Len(), len(es))
	}
}

func checkEqualEl(t *testing.T, x interface{}, want interface{}) {
	t.Helper()
	if x != want {
		t.Errorf("value = %v, want %v", x, want)
	}
}
