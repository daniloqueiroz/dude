package actions

import (
	"github.com/google/logger"
	"github.com/sahilm/fuzzy"
)


type Finder struct {
	actions map[string]Action
	symbols []string
}

func (f *Finder) All() []string {
	return f.symbols
}

func (f *Finder) Suggest(input string) []Action {
	var results []Action // an empty list

	matches := fuzzy.Find(input, f.symbols)
	for _, match := range matches {
		results = append(results, f.actions[match.Str])
	}
	logger.Infof("options for %s -> %s", input, results)
	return results
}

func (f *Finder) Do(name string) {
	logger.Infof("Selected action: %s", name)
	action := f.actions[name]
	action.Exec()
}

func FinderNew() *Finder {
	actions := make(map[string]Action)

	loadShellActions(actions)
	loadApplicationActions(actions)
	loadInternalActions(actions)
	loadPasswordsActions(actions)

	var symbols []string
	for k, _ := range actions {
		symbols = append(symbols, k)
	}
	return &Finder{
		actions: actions,
		symbols: symbols,
	}
}
