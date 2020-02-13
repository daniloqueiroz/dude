package gtk

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/gotk3/gotk3/gtk"
	"path/filepath"
	"strings"
)

func createSearch() *gtk.SearchEntry {
	search, err := gtk.SearchEntryNew()
	if err != nil {
		logger.Fatalf("Unable to create search: %v", err)
	}
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

func createScroll() *gtk.ScrolledWindow {
	vA, _ := gtk.AdjustmentNew(1, 0, 10, 1, 1, 10)
	hA, _ := gtk.AdjustmentNew(1, 1, 1, 0, 0, 1)
	scroll, _ := gtk.ScrolledWindowNew(hA, vA)
	return scroll
}

func createBox() *gtk.Box {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		logger.Fatalf("Unable to create Box:", err)
	}
	return box
}

func createPane() *gtk.Paned {
	pane, err := gtk.PanedNew(gtk.ORIENTATION_VERTICAL)
	if err != nil {
		logger.Fatalf("Unable to create Pane:", err)
	}

	return pane
}

func createStatus() *gtk.Statusbar {
	status, err := gtk.StatusbarNew()
	if err != nil {
		logger.Fatalf("Unable to create StatusBar:", err)
	}
	status.SetOpacity(0.20)
	status.SetBorderWidth(0)
	status.SetMarginTop(0)
	status.SetMarginBottom(0)
	status.SetMarginStart(0)
	status.SetMarginEnd(0)
	status.SetCanFocus(false)
	return status
}

func createWindow() *gtk.Window {
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
	return win
}

func getIconFile(category laucher.Category) string {
	var icons = map[laucher.Category]string{
		laucher.Application:       "application.png",
		laucher.Password:          "password.png",
		laucher.File:              "file.png",
		laucher.PersonalAssistant: "assistant.png",
		laucher.System:            "settings.png",
		laucher.ShellCommand:      "shell.png",
	}
	return filepath.Join(system.Config.LauncherIconsFolder, icons[category])
}

func createLabel(option laucher.ActionMeta, keyword string) *gtk.Box {
	labelWithImage, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)

	icon, err := gtk.ImageNewFromFile(getIconFile(option.Category))
	if err != nil {
		logger.Fatalf("Unable to load image: %#v", err)
	}
	labelWithImage.PackStart(icon, false, false, 5)

	label, err := gtk.LabelNew("")
	if err != nil {
		logger.Fatalf("Unable to create label: %v", err)
	}
	labelWithImage.PackStart(label, false, false, 0)

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
	return labelWithImage
}
