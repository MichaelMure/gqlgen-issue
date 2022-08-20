package lazyloading

import "gqlgen-issue/graph/model/types"

type IdentityWrapper interface {
	Id() string
	Name() (string, error)
}

// in the original code, does lazy loading and change a bit the interface (added error)
type lazyIdentity struct {
	i types.Identity
}

func (l lazyIdentity) Id() string {
	return l.i.Id
}

func (l lazyIdentity) Name() (string, error) {
	return l.i.Name, nil
}
