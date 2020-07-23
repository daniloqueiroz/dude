package plugins

import (
	"github.com/sahilm/fuzzy"
)

type Category string

const (
	Application       Category = "applications"
	Password          Category = "passwords"
	File              Category = "files"
	PersonalAssistant Category = "assistant"
	System            Category = "system"
	ShellCommand      Category = "shell"
	// ...
)

type LauncherPlugin interface {
	Category() Category
	FindActions(input string) Actions
}

// Utility functions for plugins
func FilterAction(input string, actions Actions) Actions {
	var results Actions
	matches := fuzzy.FindFrom(input, actions)
	for _, match := range matches {
		if match.Score > 0 {
			results = append(results, actions[match.Index])
		}
	}
	return results
}
