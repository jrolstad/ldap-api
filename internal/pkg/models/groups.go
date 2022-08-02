package models

type Group struct {
	Id       string
	Location string
	Type     string
	Name     string
	Members  []*User
}
