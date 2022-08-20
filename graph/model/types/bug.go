package types

type Bug struct {
	Id       string
	Comments []Comment
	Author   Identity
}

type Comment struct {
	Id     string
	Author Identity
}
