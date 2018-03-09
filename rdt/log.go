package rdt

import "github.com/gritzko/ron"

type Log struct {
	heap ron.FrameHeap
}

var LOG_UUID = ron.NewName("log")

func MakeLogReducer() ron.Reducer {
	return Log{
		heap: ron.MakeFrameHeap(ron.EventComparatorDesc, nil, 2),
	}
}

func (log Log) Features() int {
	return ron.ACID_FULL
}

func (log Log) Reduce(batch ron.Batch) ron.Frame {
	log.heap.PutAll(batch)
	spec := ron.NewSpec(
		batch[0].Type(),
		batch[0].Object(),
		batch[len(batch)-1].Event(),
		ron.ZERO_UUID,
	)
	re := ron.NewFrame()
	re.AppendStateHeader(spec)
	for !log.heap.EOF() && log.heap.Current().Event().Scheme() == ron.UUID_EVENT {
		re.AppendReduced(*log.heap.Current())
		log.heap.NextPrim()
	}
	log.heap.Clear()
	return re
}

func init() {
	ron.RDTYPES[LOG_UUID] = MakeLogReducer
}
