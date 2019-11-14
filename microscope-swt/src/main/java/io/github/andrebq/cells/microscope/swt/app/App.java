package io.github.andrebq.cells.microscope.swt.app;

import io.github.andrebq.cells.microscope.bus.Cell;
import io.github.andrebq.cells.microscope.bus.Organism;
import io.github.andrebq.cells.microscope.io.github.andrebq.cells.microscope.swt.cells.SwtCell;
import io.reactivex.rxjava3.core.Flowable;
import io.reactivex.rxjava3.disposables.Disposable;
import org.eclipse.swt.SWT;
import org.eclipse.swt.widgets.Display;
import org.eclipse.swt.widgets.Shell;
import org.eclipse.swt.widgets.Text;

import java.time.Duration;
import java.time.Instant;
import java.util.concurrent.TimeUnit;

public class App {
    public static void main(String[] string) {
        final var organism = new Organism();
        final var timeCell = organism.createCell("now", Instant.now());
        final var updateFlow = updateTime(timeCell, Duration.ofSeconds(1L));
        Display display = new Display();
        Shell shell = new Shell(display);

        Text helloWorldTest = new Text(shell, SWT.NONE);
        helloWorldTest.pack();
        SwtCell.connect(helloWorldTest, timeCell, Instant::toString);

        shell.pack();
        shell.open();
        while (!shell.isDisposed()) {
            if (!display.readAndDispatch()) display.sleep();
        }
        display.dispose();
        updateFlow.dispose();
    }

    private static Disposable updateTime(final Cell<Instant> timeCell, final Duration ofSeconds) {
        return Flowable.interval(ofSeconds.getSeconds(), TimeUnit.SECONDS)
                .subscribe((v) -> timeCell.update(Instant.now()));
    }
}
