package gtk

import (
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type GtkView struct {
	win     *gtk.Window
	pane    *LauncherPane
	control laucher.Controller
}

func NewGtkView(control laucher.Controller) laucher.View {
	gtk.Init(nil)
	pane := NewLauncherPane()
	window := createWindow()
	window.Add(pane.GetWidget())

	return &GtkView{
		win:     window,
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
	v.pane.Refresh(list)
	v.win.ShowAll()
}

func (v *GtkView) SetStatusMessage(message string) {
	//v.pane.Status.Pop(0)
	v.pane.Status.Push(0, message)
	v.pane.Status.ShowAll()
}

func (v *GtkView) ShowUI() {
	v.pane.Refresh(createListBox())
	v.win.SetFocusChild(v.pane.Search)
	v.win.Connect("destroy", func() {
		v.control.Quit(v)
	})
	v.win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		v.keyPressed(keyEvent.KeyVal())
	})
	v.pane.Search.Connect("changed", func(search *gtk.SearchEntry) {
		input, _ := search.GetText()
		v.control.InputChanged(input, v)
	})
	v.win.ShowAll()
	gtk.Main()
}
func (v *GtkView) HideUI() {
	v.win.Hide()
}
func (v *GtkView) Destroy() {
	gtk.MainQuit()
}

func (v *GtkView) keyPressed(keyVal uint) {
	if keyVal == gdk.KEY_Escape {
		v.control.Quit(v)
	} else if keyVal == gdk.KEY_Down && v.pane.Search.IsFocus() {
		v.win.SetFocusChild(v.pane.List)
	} else if keyVal == gdk.KEY_Tab {
		if v.pane.Search.IsFocus() {
			v.win.SetFocusChild(v.pane.List)
		} else {
			v.win.SetFocusChild(v.pane.Search)
		}
	} else if keyVal == gdk.KEY_Return {
		pos := 0
		if v.pane.List != nil && v.pane.List.GetSelectedRow() != nil {
			selected := v.pane.List.GetSelectedRow()
			pos = selected.GetIndex()
		}
		v.control.OptionSelected(pos, v)
		v.control.ExecuteSelected(v)
	}
}
