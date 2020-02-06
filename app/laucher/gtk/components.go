package gtk

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/gotk3/gotk3/gtk"
	"strings"
)

func createSearch() *gtk.SearchEntry {
	search, err := gtk.SearchEntryNew()
	if err != nil {
		logger.Fatalf("Unable to create search: %v", err)
	}
	//search.SetPlaceholderText(">>")
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
	win.SetSizeRequest(system.Config.LauncherWidth, system.Config.LauncherHeight)
	win.SetSkipTaskbarHint(true)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.Add(pane)

	return win
}

func createLabel(option laucher.Option, keyword string) *gtk.Label {
	label, err := gtk.LabelNew("")
	if err != nil {
		logger.Fatalf("Unable to create label: %v", err)
	}

	name := option.Name
	startIdx := strings.Index(name, keyword)
	if startIdx >= 0 {
		endIdx := startIdx + len(keyword)
		highlighted := name[:startIdx] + "<b><u>" + name[startIdx:endIdx] + "</u></b>" + name[endIdx:]
		name = highlighted
	}

	title := fmt.Sprintf(
		"<tt><big>%s</big></tt> <small>%s</small>", name, option.Description)
	label.SetMarkup(title)
	label.SetHAlign(gtk.ALIGN_START)
	return label
}
