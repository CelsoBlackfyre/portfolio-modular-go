package models

type Certification struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Issuer string `json:"issuer"`
	Date   string `json:"date"`
}
