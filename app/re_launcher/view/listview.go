package view

import (
	"github.com/daniloqueiroz/dude/app/re_launcher/plugins"
	"github.com/google/logger"
	"github.com/gotk3/gotk3/gtk"
)

type ListView struct {
	builder *gtk.Builder
}

type EventHandler func(position int)

func (lv *ListView) view() *gtk.Viewport {
	obj, err := lv.builder.GetObject(RESULT_VIEW)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}
	view, ok := obj.(*gtk.Viewport)
	if !ok {
		logger.Fatalf("Unable to load view")
	}
	return view
}

func (lv *ListView) list() *gtk.ListBox {
	obj, err := lv.builder.GetObject(RESULT_LIST)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}
	list, ok := obj.(*gtk.ListBox)
	if !ok {
		logger.Fatalf("Unable to load view")
	}

	return list
}
func wrapHandler(pos int, handler EventHandler) (func(row *gtk.ListBoxRow)) {
	return func(row *gtk.ListBoxRow) {
		handler(pos)
	}
}

func (lv *ListView) loadList(actions []plugins.Action, handler EventHandler) {
	var currentCategory plugins.Category
	pos := 0
	list := lv.list()
	list.SetSelectionMode(gtk.SELECTION_SINGLE)
	for _, action := range actions {
		if action.Category() != currentCategory {
			currentCategory = action.Category()
			list.Add(newSeparator(string(currentCategory)))
		}
		row, err := gtk.ListBoxRowNew()
		if err != nil {
			logger.Fatalf("unable to load view")
		}
		row.Add(newListEntry(action.Name(), action.Description()))
		row.Connect("activate",  wrapHandler(pos, handler))
		list.Add(row)
		pos++
	}
}

func (lv *ListView) GetView(actions []plugins.Action, handler EventHandler) *gtk.Viewport {
	lv.loadList(actions, handler)
	lv.list().SetSelectionMode(gtk.SELECTION_SINGLE)
	return lv.view()
}

// functions
func ListViewNew() *ListView {
	listBuilder, err := gtk.BuilderNew()
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}
	err = listBuilder.AddFromFile(LIST_VIEW)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}

	return &ListView{
		builder: listBuilder,
	}
}

func newListEntry(name, description string) *gtk.Box {
	builder, err := gtk.BuilderNew()
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}
	err = builder.AddFromFile(LIST_ENTRY_VIEW)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}

	obj, err := builder.GetObject(ENTRY_NAME)
	if err != nil {
		logger.Fatalf("unable to get window", err)
	}
	nameLabel, ok := obj.(*gtk.Label)
	if !ok {
		logger.Fatalf("unable to get window")
	}
	nameLabel.SetLabel(name)

	obj, err = builder.GetObject(ENTRY_DESC)
	if err != nil {
		logger.Fatalf("unable to get window", err)
	}
	descLabel, ok := obj.(*gtk.Label)
	if !ok {
		logger.Fatalf("unable to get window")
	}
	descLabel.SetLabel(description)

	obj, err = builder.GetObject(RESULT_ENTRY)
	if err != nil {
		logger.Fatalf("unable to get window", err)
	}
	box, ok := obj.(*gtk.Box)
	if !ok {
		logger.Fatalf("unable to get window")
	}

	return box
}

func newSeparator(category string) *gtk.Box {
	builder, err := gtk.BuilderNew()
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}
	err = builder.AddFromFile(CATEGORY_VIEW)
	if err != nil {
		logger.Fatalf("Unable to load view", err)
	}

	obj, err := builder.GetObject(CATEGORY_SEPARATOR)
	if err != nil {
		logger.Fatalf("unable to get window", err)
	}
	separator, ok := obj.(*gtk.Box)
	if !ok {
		logger.Fatalf("unable to get window")
	}

	obj, err = builder.GetObject(CATEGORY_NAME)
	if err != nil {
		logger.Fatalf("unable to get window", err)
	}
	label, ok := obj.(*gtk.Label)
	if !ok {
		logger.Fatalf("unable to get window")
	}
	label.SetLabel(category)

	return separator
}
