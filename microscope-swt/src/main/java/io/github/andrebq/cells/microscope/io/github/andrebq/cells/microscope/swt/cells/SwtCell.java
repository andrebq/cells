package io.github.andrebq.cells.microscope.io.github.andrebq.cells.microscope.swt.cells;

import io.github.andrebq.cells.microscope.bus.Cell;
import org.eclipse.swt.widgets.Text;

import java.util.function.Function;

public class SwtCell {
    public static <T> void connect(final Text text, final Cell<T> timeCell, final Function<T, String> format) {
        final var subscription = timeCell.watch(v -> text.getDisplay().asyncExec(() -> text.setText(format.apply(v))));
        text.addDisposeListener(e -> subscription.dispose());
    }
}
