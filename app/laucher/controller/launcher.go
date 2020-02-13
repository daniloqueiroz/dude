package controller

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/daniloqueiroz/dude/app/laucher/actions"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
)

var modes = map[rune]laucher.Category{
	':': laucher.System,
	'!': laucher.ShellCommand,
	'/': laucher.File,
	'+': laucher.PersonalAssistant,
	'@': laucher.Password,
}

const defaultStatus = "Type : for system mode, ! for shell command mode and @ for password mode"

type Launcher struct {
	finder           *actions.Loader
	availableOptions laucher.Actions
	selectedOption   int
}

func (l *Launcher) Start(v laucher.View) {
	l.finder = actions.FinderNew()
	l.availableOptions = make(laucher.Actions, 0)
	l.selectedOption = -1
	v.SetStatusMessage(defaultStatus)
	v.ShowUI()
}

func (l *Launcher) InputChanged(keyword string, view laucher.View) {
	var options []laucher.ActionMeta
	if len(keyword) == 0 {
		view.SetStatusMessage(defaultStatus)
		view.UpdateOptions(options, keyword)
		return
	}

	chars := []rune(keyword)
	var categories []laucher.Category
	category, exists := modes[chars[0]]
	if exists {
		keyword = string(chars[1:len(keyword)])
		categories = append(categories, category)
		view.SetStatusMessage(fmt.Sprintf("Active mode: %s", category))
	} else {
		for _, categoryStr := range system.Config.LauncherDefaultCategories {
			categories = append(categories, laucher.Category(categoryStr))
		}
		view.SetStatusMessage(defaultStatus)
	}

	if len(keyword) == 0 {
		// keyword len might have changed after removing category selector char
		return
	}

	l.availableOptions = make(laucher.Actions, 0)
	suggestions := l.finder.Suggest(keyword, categories...)

	for _, action := range suggestions {
		l.availableOptions = append(l.availableOptions, action)
		options = append(options, action.Details)
	}
	view.UpdateOptions(options, keyword)
}
func (l *Launcher) OptionSelected(id int, view laucher.View) {
	if len(l.availableOptions) > id {
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

func (l *Launcher) Quit(view laucher.View) {
	view.Destroy()
}
