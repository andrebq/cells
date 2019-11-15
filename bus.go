package cells

import "sync"

type (
	// bus is used by cells to manage their watchers
	bus struct {
		sync.RWMutex
		watchers map[uint64]Watcher
	}

	busSubscription struct {
		w Watcher
		b *bus
	}
)

func (bs busSubscription) Dispose() {
	bs.b.Lock()
	delete(bs.b.watchers, bs.w.uniqToken())
	bs.b.Unlock()
}

func (b *bus) init() {
	if b.watchers != nil {
		return
	}
	b.watchers = make(map[uint64]Watcher)
}

func (b *bus) add(w Watcher) Subscription {
	b.Lock()
	var bs busSubscription
	{
		bs = busSubscription{b: b}
		b.watchers[w.uniqToken()] = w
	}
	b.Unlock()

	return bs
}

func (b *bus) notify(c Cell) {
	b.RLock()
	defer b.RUnlock()

	if b.watchers == nil {
		return
	}

	for _, v := range b.watchers {
		v.Next(c)
	}
}

func (b *bus) die(c Cell) {
	b.RLock()
	defer b.RUnlock()

	if b.watchers == nil {
		return
	}

	for _, v := range b.watchers {
		v.Dead(c)
	}
}
