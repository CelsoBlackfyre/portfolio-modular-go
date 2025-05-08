package models

type Education struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	StudyProgram string `json:"studyprogram"`
	Institution  string `json:"institution"`
	Duration     string `json:"duration"`
}
