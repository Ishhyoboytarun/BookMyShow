package models

type City struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	State    string     `json:"state"`
	Theatres []*Theatre `json:"theatres"`
}
