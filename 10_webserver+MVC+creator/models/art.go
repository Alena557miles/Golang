package models

type Art struct {
	Id   int
	Name string `json:"name"`
	//Width  int    `json:"widht"`
	//Height int    `json:"height"`
	Owner string `json:"owner"`
}

func (a *Art) IsntAssigned() bool {
	return a.Owner == ""
}
