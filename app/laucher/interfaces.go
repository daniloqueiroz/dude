package laucher

// Data Model
type Category string

const (
	ShellCommand      Category = "shellCmd"
	Application       Category = "app"
	System            Category = "system"
	File		  Category = "file"
	PersonalAssistant Category = "assistant"
	// ...
)

type Option struct {
	Id          int
	Name        string
	Description string
	Category    Category
}


// View
type View interface {
	UpdateOptions(options []Option, keyword string)
	ShowUI()
	HideUI()
	Destroy()
}


// Controller - don't keep reference for the view
type Controller interface {
	Start(view View)
	InputChanged(keyword string, view View)
	SelectedOption(position int, view View)
	ExecuteSelected(view View)
	Quit(view View)
}
