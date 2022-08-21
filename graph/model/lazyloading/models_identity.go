package lazyloading

import "github.com/MichaelMure/gqlgen-issue/graph/model/types"

type IdentityWrapper interface {
	Id() string
	Name() (string, error)
}

// in the original code, does lazy loading and change a bit the interface (added error)
type LazyIdentity struct {
	i types.Identity
}

func NewLazyIdentity(i types.Identity) *LazyIdentity {
	return &LazyIdentity{i: i}
}

func (l LazyIdentity) Id() string {
	return l.i.Id
}

func (l LazyIdentity) Name() (string, error) {
	return l.i.Name, nil
}
