package models

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Pokemon        string `json:"pokemon"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
}
