package schema

type recursable interface {
	SetRecursive()
}

type recurseStack struct {
	fields  []recursable
	asStack []string
	asMap   map[string]bool
}
