package controller

import (
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/daniloqueiroz/dude/app/laucher/actions"
	"github.com/google/logger"
)

type Launcher struct {
	finder           *actions.Finder
	availableOptions []actions.Action
	selectedOption   int
}

func (l *Launcher) Start(v laucher.View) {
	l.finder = actions.FinderNew()
	l.availableOptions = make([]actions.Action, 0)
	l.selectedOption = -1
	v.ShowUI()
}

func (l *Launcher) InputChanged(keyword string, view laucher.View) {
	l.availableOptions = make([]actions.Action, 0)
	suggestions := l.finder.Suggest(keyword)
	options := make([]laucher.Option, len(suggestions))
	for idx, action := range suggestions {
		l.availableOptions = append(l.availableOptions, action)
		options[idx] = laucher.Option{
			Id:          idx,
			Name:        action.Input(),
			Description: action.Description(),
			Category:    laucher.System, // TODO fix it
		}
	}
	view.UpdateOptions(options, keyword)
}
func (l *Launcher) SelectedOption(id int, view laucher.View) {
	if len(l.availableOptions) > id  {
		l.selectedOption = id
		action := l.availableOptions[l.selectedOption]
		logger.Infof("Selected action is [%d: %v]", l.selectedOption, action)
	}
}

func (l *Launcher) ExecuteSelected(view laucher.View) {
	// TODO how to pass extra parameters to the action?
	// Maybe is already within the action
	if len(l.availableOptions) > l.selectedOption && l.selectedOption != -1 {
		view.HideUI()
		action := l.availableOptions[l.selectedOption]
		logger.Infof("Execute selected action [%d: %v]", l.selectedOption, action)
		action.Exec() // handle error
		view.Destroy()
	}
}

func (l *Launcher) Quit(view laucher.View)  {
	view.Destroy()
}
