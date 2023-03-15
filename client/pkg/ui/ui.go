package ui

import (
	"context"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UI struct {
	App  *tview.Application
	Room *Room

	Layout *tview.Grid
}

func CreateUI(cancelCtxCause context.CancelCauseFunc) *UI {
	app := tview.NewApplication()

	ui := &UI{
		App: app,
	}

	ui.Room = InitRoomsUI(ui)
	ui.Layout = ui.createGrid()

	ui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			ui.App.Stop()
			cancelCtxCause(nil)
		}
		return event
	})

	return ui
}

func (ui *UI) createGrid() *tview.Grid {
	return tview.NewGrid().SetRows(1, 0).SetColumns(0)
}

func (ui *UI) createBox() *tview.Box {
	return tview.NewBox().SetBorder(true)
}

func (ui *UI) createList() *tview.List {
	return tview.NewList()
}
