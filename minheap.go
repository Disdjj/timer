package djj_Timer

import (
	"container/heap"
	"context"
	"time"
)

type MinHeapTimer struct {
	BasicTimerConfig
	minHeaps
	ctx    context.Context
	cancel context.CancelFunc
}

func (m *MinHeapTimer) After(t time.Duration, f func()) WorkItem {
	node := &TimerNode{
		work:         f,
		scheduleTime: time.Now().Add(t),
		isLoop:       false,
		heap:         m,
	}
	m.Push(node)
	return node
}
func (m *MinHeapTimer) Schedule(t time.Duration, f func()) WorkItem {
	node := &TimerNode{
		work:         f,
		scheduleTime: time.Now().Add(t),
		isLoop:       true,
		heap:         m,
		interval:     t,
	}
	m.Push(node)
	return node
}

func (m *MinHeapTimer) Custom(next Next, f func()) WorkItem {
	node := &TimerNode{
		work:   f,
		isLoop: true,
		heap:   m,
		next:   next,
	}
	m.Push(node)
	return node
}

func (m *MinHeapTimer) Start() {
	tick := time.Tick(m.interval)
	ctx, cancel := context.WithCancel(context.TODO())
	m.ctx = ctx
	m.cancel = cancel
	for {
		select {
		case <-tick:
			now := time.Now()
			for m.Len() > 0 {
				first := m.minHeaps[0]

				if first.scheduleTime.After(now) {
					break
				}

				first.work()

				if first.isLoop {
					first.scheduleTime = first.scheduleTime.Add(first.interval)
					heap.Fix(&m.minHeaps, first.index)
				} else {
					heap.Pop(&m.minHeaps)
				}

			}
		case <-m.ctx.Done():
			m.cancel()
			return
		}
	}
}

func (m *MinHeapTimer) Stop() {
	m.cancel()
}

func (m *MinHeapTimer) removeItem(node *TimerNode) {
	if node.index < 0 || node.index > len(m.minHeaps) || len(m.minHeaps) == 0 {
		return
	}

	heap.Remove(&m.minHeaps, node.index)
}
