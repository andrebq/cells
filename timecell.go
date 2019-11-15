package cells

import (
	"sync/atomic"
	"time"
)

type (
	timeCell struct {
		baseCell
		value atomic.Value
	}

	// TimeCell contains a time value wrapped in a cell
	TimeCell interface {
		Cell
		Set(v time.Time)
		Get() time.Time
	}
)

// NewTimeCell returns a cell with a time value
func NewTimeCell(initial time.Time) TimeCell {
	tc := &timeCell{}
	tc.baseCell.init()
	tc.value.Store(initial)
	return tc
}

func (t *timeCell) Get() time.Time {
	return t.value.Load().(time.Time)
}

func (t *timeCell) Set(v time.Time) {
	t.value.Store(v)
	t.bus.notify(t)
}
