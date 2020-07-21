package view

type EventType string

const (
	SearchInputChanged EventType = "search_input_changed"
	ActionSelected     EventType = "action_selected"
	Quit               EventType = "quit"
)

type ViewEvent interface {
	Type() EventType
}

type QuitEvent struct {}
func (qe QuitEvent) Type() EventType {
	return Quit
}

type SearchEvent struct {
	Input string
}
func (se SearchEvent) Type() EventType {
	return SearchInputChanged
}

type ActionSelectedEvent struct {
	Position int
}
func (ase ActionSelectedEvent) Type() EventType {
	return ActionSelected
}