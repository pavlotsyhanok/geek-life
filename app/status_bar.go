package main

import (
	"sync/atomic"
	"time"

	"github.com/rivo/tview"
)

// StatusBar displays hints and messages at the bottom of app
type StatusBar struct {
	*tview.Pages
	message   *tview.TextView
	container *tview.Application
	token     atomic.Uint64
}

// Name of page keys
const (
	defaultPage = "default"
	messagePage = "message"
)

func prepareStatusBar(app *tview.Application) *StatusBar {
	statusBar = &StatusBar{
		Pages:     tview.NewPages(),
		message:   tview.NewTextView().SetDynamicColors(true).SetText("Loading..."),
		container: app,
	}

	statusBar.AddPage(messagePage, statusBar.message, true, true)
	statusBar.AddPage(defaultPage,
		tview.NewGrid(). // Content will not be modified, So, no need to declare explicitly
					SetColumns(0, 0, 0, 0).
					SetRows(0).
					AddItem(tview.NewTextView().SetText("Navigate List: ↓,↑ / j,k"), 0, 0, 1, 1, 0, 0, false).
					AddItem(tview.NewTextView().SetText("New Task/Project: n").SetTextAlign(tview.AlignCenter), 0, 1, 1, 1, 0, 0, false).
					AddItem(tview.NewTextView().SetText("Step back: Esc").SetTextAlign(tview.AlignCenter), 0, 2, 1, 1, 0, 0, false).
					AddItem(tview.NewTextView().SetText("Quit: Ctrl+C").SetTextAlign(tview.AlignRight), 0, 3, 1, 1, 0, 0, false),
		true,
		true,
	)

	return statusBar
}

func (bar *StatusBar) restore() {
	bar.container.QueueUpdateDraw(func() {
		bar.SwitchToPage(defaultPage)
	})
}

func (bar *StatusBar) showForSeconds(message string, timeout int) {
	if bar.container == nil {
		return
	}

	bar.message.SetText(message)
	bar.SwitchToPage(messagePage)
	token := bar.token.Add(1)

	go func() {
		time.Sleep(time.Second * time.Duration(timeout))

		if bar.token.Load() == token {
			bar.restore()
		}
	}()
}
