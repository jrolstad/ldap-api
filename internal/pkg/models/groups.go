package models

type Group struct {
	Id      string
	Domain  string
	Name    string
	Members []*User
}
