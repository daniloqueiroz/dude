package actions

import (
	"github.com/daniloqueiroz/dude/pkg/laucher"
	"github.com/google/logger"
	"github.com/sahilm/fuzzy"
)


type finder struct {
	actions map[string]laucher.Action
	symbols []string
}

func (f finder) All() []string {
	return f.symbols
}

func (f finder) Suggest(input string) []laucher.Action {
	var results []laucher.Action // an empty list

	matches := fuzzy.Find(input, f.symbols)
	for _, match := range matches {
		results = append(results, f.actions[match.Str])
	}
	logger.Infof("options for %s -> %s", input, results)
	return results
}

func (f finder) Do(name string) {
	logger.Infof("Selected action: %s", name)
	action := f.actions[name]
	action.Exec()
}

func FinderNew() laucher.Finder {
	actions := make(map[string]laucher.Action)

	loadShellActions(actions)
	loadApplicationActions(actions)
	loadInternalActions(actions)
	loadPasswordsActions(actions)

	var symbols []string
	for k, _ := range actions {
		symbols = append(symbols, k)
	}
	return finder{
		actions: actions,
		symbols: symbols,
	}
}
