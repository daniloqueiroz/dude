package laucher

import (
	"github.com/sahilm/fuzzy"
)

type Category string

const (
	Application       Category = "app"
	Password          Category = "password"
	File              Category = "file"
	PersonalAssistant Category = "assistant"
	System            Category = "system"
	ShellCommand      Category = "shellCmd"
	// ...
)

type ActionMeta struct {
	Name        string
	Description string
	Category    Category
}

type Action struct {
	Details ActionMeta
	Exec    func()
}

type Actions []Action

func (a Actions) String(i int) string {
	return a[i].Details.Name
}

func (a Actions) Len() int {
	return len(a)
}

type ActionFinder interface {
	Find(input string) Actions
}

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
