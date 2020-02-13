package gtk

import (
	"github.com/gotk3/gotk3/gtk"
)

type LauncherPane struct {
	box    *gtk.Box
	scroll *gtk.ScrolledWindow
	Search *gtk.SearchEntry
	Status *gtk.Statusbar
	List   *gtk.ListBox
}

func NewLauncherPane() *LauncherPane {
	search := createSearch()
	scroll := createScroll()
	list := createListBox()
	status := createStatus()

	box := createBox()
	box.PackStart(search, false, false, 0)
	box.PackStart(scroll, true, true, 0)
	box.PackEnd(status, false, false, 0)

	return &LauncherPane{
		box:    box,
		scroll: scroll,
		Search: search,
		Status: status,
		List:   list,
	}
}

func (l *LauncherPane) GetWidget() gtk.IWidget {
	return l.box
}

func (l *LauncherPane) Refresh(list *gtk.ListBox) {
	w, err := l.scroll.GetChild()
	if err == nil {
		w.Hide()
		l.scroll.Remove(w)
	}
	list.SetSelectionMode(gtk.SELECTION_SINGLE)
	l.scroll.Add(list)
	l.List = list
}
