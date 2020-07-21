package view

import "github.com/daniloqueiroz/dude/app/re_launcher/plugins"

type View interface {
	OnEvent(func(ev ViewEvent))
	ShowUI()
	HideUI()
	Quit()
	SetStatusMessage( string)
	ShowActions([]plugins.Action)
	ClearResults()
}