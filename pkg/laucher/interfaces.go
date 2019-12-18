package laucher

type Finder interface {
	Suggest(input string) []Action
	Do(name string)
	All() []string
}

type Action interface {
	Input() string
	Description() string
	Exec()
	String() string
}
