package models

type Group struct {
	Id      string
	Type    string
	Name    string
	Members []*User
}
