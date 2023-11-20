package djj_Timer

import "time"

type TimerNode struct {
	work         func()
	scheduleTime time.Time
	interval     time.Duration
	next         Next
	isLoop       bool
	index        int
	heap         *MinHeapTimer
}

func (t *TimerNode) Stop() {
	t.heap.removeItem(t)
}

func (t *TimerNode) Next() time.Duration {
	if t.next != nil {
		return t.next.Next()
	}
	return t.interval
}

type minHeaps []*TimerNode

func (t minHeaps) Len() int {
	return len(t)
}

func (t minHeaps) Less(i, j int) bool {
	return t[i].scheduleTime.Before(t[j].scheduleTime)
}

func (t minHeaps) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
	t[i].index = i
	t[j].index = j
}

func (t *minHeaps) Push(x any) {
	*t = append(*t, x.(*TimerNode))
	lastIndex := len(*t) - 1
	(*t)[lastIndex].index = lastIndex
}

func (t *minHeaps) Pop() any {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[0 : n-1]
	return x
}
