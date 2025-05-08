package models

type Experience struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Position string `json:"position"`
	Company  string `json:"company"`
	Tasks    string `json:"tasks"`
	Duration string `json:"duration"`
}
