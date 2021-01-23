package fifo

type NaiveListSlice []interface{}

func (l *NaiveListSlice) Len() int {
	return len(*l)
}

func (l *NaiveListSlice) PushBack(v interface{}) {
	*l = append(*l, v)
}

func (l *NaiveListSlice) PopFront() interface{} {
	list := *l
	if len(list) == 0 {
		return nil
	}
	v := list[0]
	*l = list[1:]
	return v
}

func (l *NaiveListSlice) PeekFront() interface{} {
	list := *l
	if len(list) == 0 {
		return nil
	}
	return list[0]
}
