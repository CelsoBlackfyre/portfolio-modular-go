package models

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Bio         string `json:"bio"`
	Email       string `json:"email"`
	Github      string `json:"github"`
	Linkedin    string `json:"linkedin"`
	PhoneNumber string `json:"phonenumber"`
	Address     string `json:"address"`
}
