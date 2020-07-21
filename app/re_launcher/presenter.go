package re_launcher

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/re_launcher/plugins"
	"github.com/daniloqueiroz/dude/app/re_launcher/view"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
)

var modes = map[rune]plugins.Category{
	':': plugins.System,
	'!': plugins.ShellCommand,
	'/': plugins.File,
	'+': plugins.PersonalAssistant,
	'@': plugins.Password,
}

const (
	DEFAULT_STATUS  = "Type : for system mode, ! for shell command mode and @ for password mode"
	CATEGORY_STATUS = "Active mode: %s"
)

type Presenter struct {
	view              view.View
	launcher          *Launcher
	defaultCategories []plugins.Category
}

func PresenterNew(view view.View) *Presenter {
	var categories []plugins.Category
	for _, categoryStr := range system.Config.LauncherDefaultCategories {
		categories = append(categories, plugins.Category(categoryStr))
	}
	return &Presenter{
		view:              view,
		launcher:          LauncherNew(),
		defaultCategories: categories,
	}
}

func (p *Presenter) onEvent(viewEvent interface{}) {
	logger.Infof("Event %s received", viewEvent)
	switch ev := viewEvent.(type) {
	case view.QuitEvent:
		p.view.Quit()
	case view.SearchChangedEvent:
		p.onSearchInputChanged(ev.Input)
	case view.ActionSelectedEvent:
		// TODO get result and process
		p.launcher.ExecuteOption(ev.Position)
		// TODO handle the result
		p.view.Quit()
	}
}

func (p *Presenter) onSearchInputChanged(keyword string) {
	// TODO check mode - sub menu
	if len(keyword) == 0 {
		p.view.SetStatusMessage(DEFAULT_STATUS)
		p.view.ClearResults()
		return
	}

	var status_msg string
	var categories []plugins.Category
	chars := []rune(keyword)
	category, exists := modes[chars[0]]
	if exists {
		// Filter by category
		keyword = string(chars[1:len(keyword)])
		categories = append(categories, category)
		status_msg = fmt.Sprintf(CATEGORY_STATUS, category)
	} else {
		// No category filter
		categories = p.defaultCategories
		status_msg = DEFAULT_STATUS
	}
	p.view.SetStatusMessage(status_msg)

	if len(keyword) == 0 {
		// keyword len might have changed after removing category selector char
		p.view.ClearResults()
		return
	}

	p.launcher.RefreshOptions(keyword, categories)
	p.view.ShowActions(p.launcher.AvailableActions())
}

func (p *Presenter) Init() {
	p.view.OnEvent(p.onEvent)
	p.view.SetStatusMessage(DEFAULT_STATUS)
	p.view.ShowUI()
}
