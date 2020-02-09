package gtk

import (
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type GtkView struct {
	win     *gtk.Window
	search  *gtk.SearchEntry
	pane    *gtk.Paned
	list    *gtk.ListBox
	control laucher.Controller
}

func NewGtkView(control laucher.Controller) laucher.View {
	gtk.Init(nil)
	search := createSearch()
	pane := createPane(search)
	window := createWindow(pane)

	return &GtkView{
		win:     window,
		search:  search,
		pane:    pane,
		control: control,
	}
}

func (v *GtkView) UpdateOptions(options []laucher.ActionMeta, keyword string) {
	list := createListBox()
	for _, option := range options {
		label := createLabel(option, keyword)
		list.Add(label)
	}
	v.refresh(list)
}

func (v *GtkView) ShowUI() {
	v.refresh(createListBox())
	v.win.SetFocusChild(v.search)
	v.win.Connect("destroy", func() {
		v.control.Quit(v)
	})
	v.win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		v.keyPressed(keyEvent.KeyVal())
	})
	v.search.Connect("changed", func(search *gtk.SearchEntry) {
		input, _ := search.GetText()
		v.control.InputChanged(input, v)
	})
	gtk.Main()
}
func (v *GtkView) HideUI() {
	v.win.Hide()
}
func (v *GtkView) Destroy() {
	gtk.MainQuit()
}

func (v *GtkView) refresh(list *gtk.ListBox) {
	w, err := v.pane.GetChild2()
	if err == nil {
		w.Hide()
		v.pane.Remove(w)
	}
	v.list = list
	v.list.SetSelectionMode(gtk.SELECTION_SINGLE)
	v.pane.Add2(createScroll(v.list))
	v.win.ShowAll()
}

func (v *GtkView) keyPressed(keyVal uint) {
	if keyVal == gdk.KEY_Escape {
		v.control.Quit(v)
	} else if keyVal == gdk.KEY_Down && v.search.IsFocus() {
		v.win.SetFocusChild(v.list)
	} else if keyVal == gdk.KEY_Tab {
		if v.search.IsFocus() {
			v.win.SetFocusChild(v.list)
		} else {
			v.win.SetFocusChild(v.search)
		}
	} else if keyVal == gdk.KEY_Return {
		pos := 0
		if v.list != nil && v.list.GetSelectedRow() != nil {
			selected := v.list.GetSelectedRow()
			pos = selected.GetIndex()
		}
		v.control.OptionSelected(pos, v)
		v.control.ExecuteSelected(v)
	}
}
