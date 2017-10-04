package RON

func NewBufferIterator(data []byte) (i Iterator) {
	i.state.data = data
	i.Parse()
	return
}

func NewStringIterator(data string) (i Iterator) {
	return NewBufferIterator([]byte(data))
}

func (i Iterator) IsLast() bool {
	return i.state.p >= len(i.state.data)
}

func (i *Iterator) Next() Op {
	i.Parse()
	return i.Op
}
