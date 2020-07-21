package view

import (
	"github.com/daniloqueiroz/dude/app/re_launcher/plugins"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type MainView struct {
	builder *gtk.Builder
	EvChan  chan ViewEvent
}

func (v *MainView) scrollWindow() *gtk.ScrolledWindow {
	obj, err := v.builder.GetObject(RESULT_WINDOW)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}
	win, ok := obj.(*gtk.ScrolledWindow)
	if !ok {
		logger.Fatalf("Unable to load view")
	}
	return win
}

func (v *MainView) window() *gtk.Window {
	obj, err := v.builder.GetObject(WINDOW_NAME)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}
	window, ok := obj.(*gtk.Window)
	if !ok {
		logger.Fatalf("Unable to load view")
	}
	return window
}

func (v *MainView) search() *gtk.SearchEntry {
	obj, err := v.builder.GetObject(SEARCH_NAME)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}
	search, ok := obj.(*gtk.SearchEntry)
	if !ok {
		logger.Fatalf("Unable to load view")
	}
	return search
}

func (v *MainView) status() *gtk.Statusbar {
	obj, err := v.builder.GetObject(STATUS_NAME)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}
	status, ok := obj.(*gtk.Statusbar)
	if !ok {
		logger.Fatalf("Unable to load view")
	}
	return status
}

func (v *MainView) keyPressed(keyVal uint) {
	search := v.search()
	if keyVal == gdk.KEY_Escape {
		v.EvChan <- QuitEvent{}
	} else if keyVal == gdk.KEY_Return && search.HasFocus() {
		v.listEntrySelected(0)
	}
}

func (v MainView) searchBarChanged(input string) {
	v.EvChan <- SearchEvent{
		Input: input,
	}
}

func (v MainView) listEntrySelected(position int) {
	v.EvChan <- ActionSelectedEvent{
		Position: position,
	}
}

func (v MainView) OnEvent(handler func(ev ViewEvent)) {
	go func() {
		for ev := range v.EvChan {
			handler(ev)
		}
	}()
}

func (v MainView) ShowUI() {
	stock, _ := gtk.CssProviderNew()
	if err := stock.LoadFromPath(LAUNCHER_CSS); err != nil {
		logger.Fatalf("Failed to load CSS", err)
	}
	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, stock, gtk.STYLE_PROVIDER_PRIORITY_USER)

	signals := map[string]interface{}{
		"launcherQuit": func() {
			v.EvChan <- QuitEvent{}
		},
		"launcherKeyPressed": func(w *gtk.Window, ev *gdk.Event) {
			keyEvent := &gdk.EventKey{ev}
			keyVal := keyEvent.KeyVal()
			v.keyPressed(keyVal)
		},
		"searchBarChanged": func(searchInput *gtk.SearchEntry) {
			input, _ := searchInput.GetText()
			v.searchBarChanged(input)
		},
	}
	v.builder.ConnectSignals(signals)

	win := v.window()

	win.SetSizeRequest(system.Config.LauncherWidth, system.Config.LauncherHeight)
	win.ShowAll()
	gtk.Main()
}

func (v MainView) HideUI() {
	v.window().Hide()
}
func (v MainView) Quit() {
	close(v.EvChan)
	gtk.MainQuit()
}

func (v MainView) SetStatusMessage(message string) {
	sb := v.status()
	glib.IdleAdd(func() {
		sb.Push(0, message)
		sb.ShowAll()
	})
}

func (v MainView) ShowActions(actions []plugins.Action) {
	sw := v.scrollWindow()
	v.ClearResults()

	glib.IdleAdd(func() {
		sw.Add(ListViewNew().GetView(actions, v.listEntrySelected))
		sw.ShowAll()
	})

}

func (v MainView) ClearResults() {
	sw := v.scrollWindow()
	child, err := sw.GetChild()
	if err == nil {
		glib.IdleAdd(func() {
			child.Hide()
			sw.Remove(child)
		})
	}
}

func ViewNew() View {
	gtk.Init(nil)
	builder, err := gtk.BuilderNew()
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}

	err = builder.AddFromFile(LAUNCHER_UI)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}

	return MainView{
		builder: builder,
		EvChan:  make(chan ViewEvent),
	}
}
