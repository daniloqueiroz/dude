package view

type QuitEvent struct{}

type SearchChangedEvent struct {
	Input string
}

type ActionSelectedEvent struct {
	Position int
}
