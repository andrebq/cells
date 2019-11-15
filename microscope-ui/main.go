package main

import (
	"time"

	"github.com/andlabs/ui"
	"github.com/andrebq/cells"

	// add window manifest
	_ "github.com/andlabs/ui/winmanifest"
)

func onMain(action func()) func() {
	return func() { ui.QueueMain(action) }
}

func setupUI() {
	timeCell := cells.NewTimeCell(time.Now())
	go func() {
		for range time.NewTicker(time.Second).C {
			timeCell.Set(time.Now())
		}
	}()

	var rootWin *ui.Window
	rootWin = ui.NewWindow("Microscope", 480, 240, true)
	rootWin.OnClosing(func(w *ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		rootWin.Destroy()
		return true
	})

	vbox := ui.NewVerticalBox()

	vbox.SetPadded(true)
	vbox.Append(ui.NewButton("Button"), false)
	vbox.Append(ui.NewButton("Checkbox"), false)
	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)
	vbox.Append(ui.NewHorizontalSeparator(), false)

	clockbtn := ui.NewButton("Clock button")

	timeCell.Observe(cells.Act(
		onMain(func() {
			clockbtn.SetText(
				timeCell.Get().Format(time.RFC3339Nano))
		})))

	vbox.Append(clockbtn, false)

	rootWin.SetChild(vbox)

	rootWin.Show()
}

func main() {
	ui.Main(setupUI)
}
