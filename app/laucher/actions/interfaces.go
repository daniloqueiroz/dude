package actions

type Action interface {
	Input() string
	Description() string
	Exec()
	String() string
}
