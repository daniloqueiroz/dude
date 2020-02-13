package actions

import (
	"github.com/daniloqueiroz/dude/app/laucher"
)

type Loader struct {
	finders map[laucher.Category]laucher.ActionFinder
}

func (f *Loader) Suggest(input string, categories ...laucher.Category) laucher.Actions {
	var options laucher.Actions

	for _, name := range categories {
		finder := f.finders[name]
		suggested := finder.Find(input)
		options = append(options, suggested...)
	}

	return options
}

func FinderNew() *Loader {
	return &Loader{
		finders: map[laucher.Category]laucher.ActionFinder{
			laucher.Application:  &Applications{},
			laucher.Password:     &Pass{},
			laucher.System:       &Internal{},
			laucher.ShellCommand: &Shell{},
		},
	}
}
