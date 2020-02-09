package actions

import (
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/google/logger"
	"reflect"
)

type Loader struct {
	finders  map[laucher.Category]laucher.ActionFinder
	prefixes map[rune]laucher.Category
}

func (f *Loader) Suggest(input string) laucher.Actions {
	var options laucher.Actions

	chars := []rune(input)
	category, exists := f.prefixes[chars[0]]
	if exists && len(chars) >= 2 {
		finder := f.finders[category]
		logger.Infof("Category filter: %s - Using Finder -> %s", category, finderName(finder))
		input = string(chars[1:len(input)])
		suggested := finder.Find(input)
		options = append(options, suggested...)
	} else {
		logger.Infof("All Categories")
		for _, finder := range f.finders {
			logger.Infof("Using Finder -> %s", finderName(finder))
			suggested := finder.Find(input)
			options = append(options, suggested...)
		}
	}

	return options
}

func finderName(finder laucher.ActionFinder) string {
	return reflect.TypeOf(finder).String()
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
