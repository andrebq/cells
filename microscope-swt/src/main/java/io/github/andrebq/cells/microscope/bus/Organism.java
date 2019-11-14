package io.github.andrebq.cells.microscope.bus;

import java.util.Map;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;

public final class Organism {
    private final Map<Object, Cell<?>> cellMap = new ConcurrentHashMap<>();

    public Organism() {
    }

    public final <T> Cell<T> createCell(final String cellid, final T value) {
        final var cell = new Cell<T>(cellid, value);
        cellMap.putIfAbsent(cell.getKey(), cell);
        return cell;
    }

    public final <T> Optional<Cell<T>> getCell(final String cellid) {
        return Optional.ofNullable((Cell<T>) cellMap.get(cellid));
    }
}
