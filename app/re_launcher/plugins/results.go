package plugins

type Result interface {
}

type Empty struct {
}

type SubActions struct {
	SubActions Actions
}
