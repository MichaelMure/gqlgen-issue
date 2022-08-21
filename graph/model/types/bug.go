package types

import (
	"github.com/MichaelMure/gqlgen-issue/graph/model/lazyloading"
)

type Bug struct {
	Id       string
	Comments []Comment
	Author   Identity
}

type Comment struct {
	Id     string
	Author Identity
}

func (c Comment) IsAuthored() {}

func (c Comment) GetAuthor() lazyloading.IdentityWrapper {
	/*

		ISSUE #1 : implementing the generated interface create a cyclic dependency. It's also really wrong
		          for the lower level to know about the upper level.

	*/
	return lazyloading.NewLazyIdentity(c.Author)
}
