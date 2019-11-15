package cells

type (
	// Cell contains a unit of infromation that can notify others about changes to its internal state
	Cell interface {
		Observe(Watcher) Subscription
	}

	// Watcher is by cells to notify about changes in their internal state
	Watcher interface {
		// Next is called with the cell with the new state configured
		Next(c Cell)

		// Dead is called when the cell dies
		Dead(c Cell)

		uniqToken() uint64
	}

	// Subscription is bound to a Cell and is only valid while the cell is alive
	Subscription interface {
		// Dispose is called to unsubscribe from a cell
		Dispose()
	}

	baseCell struct {
		bus bus
	}
)

func (b *baseCell) init() {
	b.bus.init()
}

func (b *baseCell) Observe(w Watcher) Subscription {
	if w == nil {
		return nil
	}
	return b.bus.add(w)
}
