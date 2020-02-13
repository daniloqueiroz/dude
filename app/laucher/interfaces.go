package laucher

// View
type View interface {
	UpdateOptions(options []ActionMeta, keyword string)
	SetStatusMessage(message string)
	ShowUI()
	HideUI()
	Destroy()
}

// Controller - don't keep reference for the view
type Controller interface {
	Start(view View)
	InputChanged(keyword string, view View)
	OptionSelected(position int, view View)
	ExecuteSelected(view View)
	Quit(view View)
}
