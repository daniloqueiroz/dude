package re_launcher

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/re_launcher/plugins"
	"github.com/daniloqueiroz/dude/app/re_launcher/view"
	"github.com/daniloqueiroz/dude/app/system"
)

var filter = map[rune]plugins.Category{
	':': plugins.System,
	'!': plugins.ShellCommand,
	'@': plugins.Password,
	'%': plugins.Web,
}

const (
	DEFAULT_STATUS  = "Type : for system mode, ! for shell command mode, @ for password mode and % for web mode"
	CATEGORY_STATUS = "Active mode: %s"
)

type Mode string

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

func (p *Presenter) processResult(result plugins.Result) {
	switch result.(type) {
	case *plugins.SubActions:
		p.view.ClearSearch()
		p.view.SetStatusMessage(fmt.Sprintf(CATEGORY_STATUS, p.launcher.GetSelectedCategories()[0]))
		p.view.ShowActions(p.launcher.AvailableActions())
	default:
		p.view.Quit()
	}

}

func (p *Presenter) onEvent(viewEvent interface{}) {
	switch ev := viewEvent.(type) {
	case view.QuitEvent:
		p.view.Quit()
	case view.SearchChangedEvent:
		p.onSearchInputChanged(ev.Input)
	case view.ActionSelectedEvent:
		result := p.launcher.ExecuteOption(ev.Position)
		p.processResult(result)
	}
}

func (p *Presenter) onSearchInputChanged(keyword string) {
	if p.launcher.GetMode() == MainMenu {
		p.launcher.SelectCategories(p.defaultCategories)
		p.view.SetStatusMessage(DEFAULT_STATUS)

		if len(keyword) == 0 {
			p.view.ClearResults()
			return
		} else {
			chars := []rune(keyword)
			category, exists := filter[chars[0]]
			if exists {
				// Filter by category
				keyword = string(chars[1:len(keyword)])
				p.launcher.SelectCategories([]plugins.Category{category})
				p.view.SetStatusMessage(fmt.Sprintf(CATEGORY_STATUS, category))
			}
		}
	}

	if len(keyword) == 0 {
		// keyword len might have changed after removing category selector char
		p.view.ClearResults()
		return
	}

	p.launcher.RefreshOptions(keyword)
	p.view.ShowActions(p.launcher.AvailableActions())
}

func (p *Presenter) Init() {
	p.view.OnEvent(p.onEvent)
	p.view.SetStatusMessage(DEFAULT_STATUS)
	p.view.ShowUI()
}
