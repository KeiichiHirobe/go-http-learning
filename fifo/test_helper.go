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
