package lazyloading

import "gqlgen-issue/graph/model/types"

type BugWrapper interface {
	Id() string
	Comments() ([]types.Comment, error)
	Author() (IdentityWrapper, error)

	IsAuthored()
}

// in the original code, does lazy loading and change a bit the interface (added error)
type lazyBug struct {
	b types.Bug
}

func (l lazyBug) Id() string {
	return l.b.Id
}

func (l lazyBug) Comments() ([]types.Comment, error) {
	return l.b.Comments, nil
}

func (l lazyBug) Author() (IdentityWrapper, error) {
	return lazyIdentity{i: l.b.Author}, nil
}

func (l lazyBug) IsAuthored() {}
