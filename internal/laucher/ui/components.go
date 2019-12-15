package ui

import (
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/google/logger"
	"github.com/gotk3/gotk3/gtk"
)

func createSearch() *gtk.SearchEntry {
	search, err := gtk.SearchEntryNew()
	if err != nil {
		logger.Fatalf("Unable to create search: %v", err)
	}
	search.SetPlaceholderText(">>")
	search.SetHExpand(true)
	search.SetActivatesDefault(true)
	return search
}

func createListBox() *gtk.ListBox {
	list, err := gtk.ListBoxNew()
	if err != nil {
		logger.Fatalf("Unable to create ListBox: %v", err)
	}
	return list
}

func createScroll(list *gtk.ListBox) *gtk.ScrolledWindow {
	vA, _ := gtk.AdjustmentNew(1, 0, 10, 1, 1, 10)
	hA, _ := gtk.AdjustmentNew(1, 1, 1, 0, 0, 1)
	scroll, _ := gtk.ScrolledWindowNew(hA, vA)
	scroll.Add(list)
	return scroll
}

func createPane(search *gtk.SearchEntry) *gtk.Paned {
	pane, err := gtk.PanedNew(gtk.ORIENTATION_VERTICAL)
	if err != nil {
		logger.Fatalf("Unable to create Pane:", err)
	}

	pane.Add(search)
	return pane
}

func createWindow(pane *gtk.Paned) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		logger.Fatalf("Unable to create window: %v", err)
	}
	win.SetTitle("Launcher")

	if err != nil {
		logger.Fatalf("Unable to create Window: %v", err)
	}
	win.SetDecorated(false)
	win.SetSizeRequest(commons.Config.LauncherWidth, commons.Config.LauncherHeight)
	win.SetSkipTaskbarHint(true)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.Add(pane)

	return win
}
