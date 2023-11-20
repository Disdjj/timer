package djj_Timer

import (
	"time"
)

type Next interface {
	Next() time.Duration
}

type WorkItem interface {
	Stop()
}
type Timer interface {
	After(t time.Duration, f func()) WorkItem
	Schedule(t time.Duration, f func()) WorkItem
	Custom(next Next, f func()) WorkItem
	Start()
	Stop()
}

type BasicTimerConfig struct {
	interval time.Duration
}

type Option func(config *BasicTimerConfig) *BasicTimerConfig

func WithInterval(interval time.Duration) Option {
	return func(config *BasicTimerConfig) *BasicTimerConfig {
		config.interval = interval
		return config
	}
}
func NewTimer(options ...Option) Timer {
	t := &MinHeapTimer{}
	for _, o := range options {
		o(&t.BasicTimerConfig)
	}
	return t
}
