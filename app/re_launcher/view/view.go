package view

import "github.com/daniloqueiroz/dude/app/re_launcher/plugins"

type View interface {
	OnEvent(func(interface{}))
	ShowUI()
	HideUI()
	Quit()
	SetStatusMessage(string)
	ShowActions([]plugins.Action)
	ClearResults()
	ClearSearch()
	SetSearch(search string)
}
