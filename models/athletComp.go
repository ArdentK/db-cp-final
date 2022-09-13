package models

type AthletComp struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	IDComp int    `json:"id_competition"`
	Status bool
}
