package cells

import "sync/atomic"

type (
	watcher struct {
		id     uint64
		nexter func(c Cell)
		deader func(c Cell)
	}
)

var (
	watcherid uint64
)

// Act returns a watcher to execute action when the cell changes
func Act(action func()) Watcher {
	return watcher{
		id:     nextWatcherID(),
		nexter: func(_ Cell) { action() },
		deader: nil,
	}
}

func nextWatcherID() uint64 {
	return atomic.AddUint64(&watcherid, 1)
}

func (w watcher) uniqToken() uint64 {
	return w.id
}

func (w watcher) Next(c Cell) {
	if w.nexter == nil {
		return
	}
	w.nexter(c)
}

func (w watcher) Dead(c Cell) {
	deader := w.deader

	w.nexter = nil
	w.deader = nil

	if deader == nil {
		return
	}
	deader(c)
}
