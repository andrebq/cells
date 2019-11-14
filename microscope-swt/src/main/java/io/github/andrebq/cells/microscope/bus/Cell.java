package io.github.andrebq.cells.microscope.bus;

import io.reactivex.rxjava3.disposables.Disposable;
import io.reactivex.rxjava3.subjects.PublishSubject;

import java.util.function.Consumer;

public final class Cell<T> {
    private volatile T v;
    private final String id;
    private final PublishSubject<T> subject;

    Cell(final String cellid, T initialValue) {
        id = cellid;
        subject = PublishSubject.create();
        v = initialValue;
    }

    public String getKey() {
        return id;
    }

    public Disposable watch(Consumer<T> consumer) {
        return subject.subscribe(consumer::accept);
    }

    public T update(T newValue) {
        final var old = v;
        v = newValue;
        subject.onNext(newValue);
        return old;
    }
}
