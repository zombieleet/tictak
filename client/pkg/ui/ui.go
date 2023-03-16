package ui

import (
	"context"
	_ "fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UI struct {
	App  *tview.Application
	Room *Room

	MainLayout    *tview.Grid
	ContentLayout *tview.Grid
}

func CreateUI(cancelCtxCause context.CancelCauseFunc) *UI {
	app := tview.NewApplication()

	ui := &UI{
		App: app,
	}

	ui.Room = InitRoomsUI(ui)

	ui.createMainLayout()

	ui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			ui.App.Stop()
			cancelCtxCause(nil)
		}
		return event
	})

	return ui
}

// createMainLayout creates the main layout of the program
// 1. game title text
// 2. content layout (it appends server information and how long connection has been established with the server)
func (ui *UI) createMainLayout() {

	ui.MainLayout = ui.createGrid([]int{1}).SetBorders(true)
	ui.ContentLayout = ui.createGrid([]int{1}).SetBorders(true)

	ui.MainLayout.AddItem(
		ui.createText("TICKTATOE GAME OVER TCP", tview.AlignCenter),
		0, 0, 1, 1, 0, 0,
		false,
	)

	//ui.ContentLayout.AddItem(nil, 0, 1, 2, 1, 0, 0, false)
	ui.MainLayout.AddItem(ui.ContentLayout, 1, 0, 1, 1, 0, 0, false)
}

func (ui *UI) createGrid(rows []int) *tview.Grid {
	return tview.NewGrid().SetRows(rows...).SetColumns(0)
}

func (ui *UI) createBox() *tview.Box {
	return tview.NewBox()
}

func (ui *UI) createList() *tview.List {
	return tview.NewList()
}

func (ui *UI) createText(text string, align int) *tview.TextView {
	return tview.NewTextView().SetText(text).SetTextAlign(align)
}
