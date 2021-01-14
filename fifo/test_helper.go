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
