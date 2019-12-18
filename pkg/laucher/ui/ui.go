package ui

import (
	"fmt"
	"github.com/daniloqueiroz/dude/pkg/laucher"
	"github.com/daniloqueiroz/dude/pkg/laucher/actions"
	"github.com/google/logger"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"strings"
)

type UI struct {
	win    *gtk.Window
	search *gtk.SearchEntry
	pane   *gtk.Paned
	list   *gtk.ListBox
	items  []laucher.Action
	finder laucher.Finder
}

func NewUI() UI {
	gtk.Init(nil)
	search := createSearch()
	pane := createPane(search)
	window := createWindow(pane)
	finder := actions.FinderNew()

	ui := UI{
		win:    window,
		search: search,
		pane:   pane,
		list:   nil,
		finder: finder,
		items:  nil,
	}

	return ui
}

func (ui *UI) Show() {
	ui.refresh(createListBox())
	ui.win.SetFocusChild(ui.search)
	ui.win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	ui.win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		ui.keyPressed(keyEvent.KeyVal())
	})
	ui.search.Connect("changed", func(search *gtk.SearchEntry) {
		content, _ := search.GetText()
		ui.searchUpdated(content)
	})
	gtk.Main()
}

func (ui *UI) keyPressed(keyVal uint) {
	if keyVal == gdk.KEY_Escape {
		gtk.MainQuit()
	} else if keyVal == gdk.KEY_Return {
		var action laucher.Action = nil
		if ui.items != nil && len(ui.items) == 1 {
			ui.win.Hide()
			action = ui.items[0].(laucher.Action)
		} else if ui.list != nil && ui.list.GetSelectedRow() != nil{
			ui.win.Hide()
			selected := ui.list.GetSelectedRow()
			action = ui.items[selected.GetIndex()]
		}
		if action != nil {
			logger.Infof("Action selected: %s", action.Input())
			action.Exec()
			gtk.MainQuit()
		} else {
			logger.Info("No action selected")
		}
	}
}

func (ui *UI) searchUpdated(keyword string) {
	suggestions := ui.finder.Suggest(keyword)
	list := createListBox()
	ui.items = make([]laucher.Action, 0)
	for _, action := range suggestions {
		label := ui.createLabel(action, keyword)
		ui.items = append(ui.items, action)
		list.Add(label)
	}
	ui.refresh(list)
}

func (ui *UI) createLabel(action laucher.Action, keyword string) *gtk.Label {
	label, err := gtk.LabelNew("")
	if err != nil {
		logger.Fatalf("Unable to create label: %v", err)
	}

	name := action.Input()
	startIdx := strings.Index(name, keyword)
	if startIdx >= 0 {
		endIdx := startIdx + len(keyword)
		highlighted := name[:startIdx] + "<b><u>" + name[startIdx:endIdx] + "</u></b>" + name[endIdx:]
		name = highlighted
	}

	title := fmt.Sprintf(
		"<tt><big>%s</big></tt> <small>%s</small>", name, action.Description())
	label.SetMarkup(title)
	label.SetHAlign(gtk.ALIGN_START)
	return label
}

func (ui *UI) refresh(list *gtk.ListBox) {
	w, err := ui.pane.GetChild2()
	if err == nil {
		w.Hide()
		ui.pane.Remove(w)
	}
	ui.list = list
	ui.list.SetSelectionMode(gtk.SELECTION_SINGLE)
	ui.pane.Add2(createScroll(ui.list))
	ui.win.ShowAll()
}
