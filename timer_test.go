package djj_Timer

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	timer := NewTimer(WithInterval(100 * time.Millisecond))
	timer.Schedule(
		1*time.Second, func() {
			t.Log("after 1 second")
		},
	)
	go timer.Start()

	time.Sleep(10 * time.Second)
	timer.Stop()
}
