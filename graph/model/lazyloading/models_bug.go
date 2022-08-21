package lazyloading

import (
	"github.com/MichaelMure/gqlgen-issue/graph/model"
	"github.com/MichaelMure/gqlgen-issue/graph/model/types"
)

type BugWrapper interface {
	Id() string
	Comments() ([]types.Comment, error)
	Author() (IdentityWrapper, error)

	/*

		ISSUE #3: embedding the interface is not recognized as implementing it.
				Copying the functions does work though.

	*/
	// model.Authored

	IsAuthored()
	// The author of this object.
	GetAuthor() IdentityWrapper
}

var _ model.Authored = &LazyBug{}

// in the original code, does lazy loading and change a bit the interface (added error)
type LazyBug struct {
	b types.Bug
}

func NewLazyBug(b types.Bug) *LazyBug {
	return &LazyBug{b: b}
}

func (l LazyBug) Id() string {
	return l.b.Id
}

func (l LazyBug) Comments() ([]types.Comment, error) {
	return l.b.Comments, nil
}

func (l LazyBug) Author() (IdentityWrapper, error) {
	return LazyIdentity{i: l.b.Author}, nil
}

func (l LazyBug) IsAuthored() {}

func (l LazyBug) GetAuthor() IdentityWrapper {
	author, err := l.Author()
	if err != nil {
		/*

			ISSUE #2: Should not panic, but can't return an error!
					(yes in this toy example I don't really have an error, but in the real thing I have lazy
					loading that *can* fail)

		*/
		panic(err)
	}
	return author
}
