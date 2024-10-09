package model_repository

type Person struct {
	Id    string
	Name  string
	Image string
	// why using isPrivate
	IsPrivate bool
}
