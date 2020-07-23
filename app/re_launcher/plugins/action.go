package plugins

type Action interface {
	Category() Category
	Name() string
	Description() string
	Execute() Result // TODO different result types (Empty, SubMenu, Data)
}

type Actions []Action

func (a Actions) String(i int) string {
	return a[i].Name()
}

func (a Actions) Len() int {
	return len(a)
}
