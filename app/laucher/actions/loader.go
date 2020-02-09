package actions

import (
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/daniloqueiroz/dude/app/system"
)

type Loader struct {
	finders  map[laucher.Category]laucher.ActionFinder
	prefixes map[rune]laucher.Category
}

func (f *Loader) Suggest(input string) laucher.Actions {
	var options laucher.Actions

	if len(input) == 0 {
		return options
	}

	chars := []rune(input)
	category, exists := f.prefixes[chars[0]]
	if exists && len(chars) >= 2 {
		finder := f.finders[category]
		input = string(chars[1:len(input)])
		suggested := finder.Find(input)
		options = append(options, suggested...)
	} else {
		for _, name := range system.Config.LauncherDefaultFinders {
			finder := f.finders[laucher.Category(name)]
			suggested := finder.Find(input)
			options = append(options, suggested...)
		}
	}

	return options
}

func FinderNew() *Loader {
	return &Loader{
		prefixes: map[rune]laucher.Category{
			':': laucher.System,
			'!': laucher.ShellCommand,
			'/': laucher.File,
			'+': laucher.PersonalAssistant,
			'@': laucher.Password,
		},
		finders: map[laucher.Category]laucher.ActionFinder{
			laucher.Application:  &Applications{},
			laucher.Password:     &Pass{},
			laucher.System:       &Internal{},
			laucher.ShellCommand: &Shell{},
		},
	}
}
